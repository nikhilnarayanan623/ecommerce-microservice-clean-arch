package usecase

import (
	"context"
	"errors"
	"fmt"

	clientInterface "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/domain"
	repo "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/response"
)

type orderUseCase struct {
	repo          repo.OrderRepository
	cartClient    clientInterface.CartClient
	productClient clientInterface.ProductClient
}

func NewOrderUseCase(repo repo.OrderRepository, cartClient clientInterface.CartClient,
	productClient clientInterface.ProductClient) interfaces.OrderUseCase {
	return &orderUseCase{
		repo:          repo,
		cartClient:    cartClient,
		productClient: productClient,
	}
}

var (
	ErrCartIsEmpty            = errors.New("user cart is empty")
	ErrCartIsNotValidForOrder = errors.New("user cart is not valid for order cart items qty not met with product qty_in_stock")
)

func (c *orderUseCase) PlaceShopOrder(ctx context.Context, userID uint64) (uint64, error) {

	cart, err := c.cartClient.FindCart(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to find user cart \nerror:%w", err)
	}
	if cart.TotalPrice == 0 {
		return 0, ErrCartIsEmpty
	}

	// check cart items are valid for order
	if ok := c.isCartIsValidForOrder(cart); !ok {
		return 0, ErrCartIsNotValidForOrder
	}

	//Transaction for order
	var shopOrderID uint64
	err = c.repo.Transaction(func(trxRepo repo.OrderRepository) error {
		// save the shop order
		shopOrderID, err = trxRepo.SaveShopOrder(ctx, domain.ShopOrder{
			UserID:          userID,
			OrderTotalPrice: cart.TotalPrice,
		})
		if err != nil {
			return fmt.Errorf("failed to save shop_order \nerror:%w", err)
		}
		// create a stock decreasing array for stock decrease
		stocksToDecrease := make([]request.StockDecrease, len(cart.CartItems))

		// create multiple order lines
		for i, cartItem := range cart.CartItems {
			// save order lines one by one
			err := trxRepo.SaveOrderLine(ctx, domain.OrderLine{
				ProductItemID: cartItem.ProductItemID,
				ShopOrderID:   shopOrderID,
				Qty:           cartItem.Qty,
				Price:         cartItem.Price,
			})
			if err != nil {
				return fmt.Errorf("failed to save order line \nerror:%w", err)
			}

			// add the product to decrease stock
			stocksToDecrease[i].SKU = cartItem.SKU
			stocksToDecrease[i].QtyToDecrease = cartItem.Qty
		}
		// Remove all cart items
		err = c.cartClient.RemoveAllCartItems(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to remove cart items \nerror:%w", err)
		}

		// decrease all products quantity
		err := c.productClient.DecreaseMultipleStocks(ctx, stocksToDecrease)
		if err != nil {
			return fmt.Errorf("failed to decrease product quantity \nerror:%w", err)
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return shopOrderID, nil

}

func (c *orderUseCase) isCartIsValidForOrder(cart response.Cart) (valid bool) {

	for _, cartItem := range cart.CartItems {
		if cartItem.Qty > cartItem.QtyInStock {
			return false
		}
	}
	return true
}

func (c *orderUseCase) FindAllShopOrders(ctx context.Context, userID uint64, pagination request.Pagination) ([]response.ShopOrder, error) {

	shopOrders, err := c.repo.FindAllShopOrdersByUserID(ctx, userID, pagination)

	return shopOrders, err
}

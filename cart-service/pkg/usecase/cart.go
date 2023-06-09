package usecase

import (
	"context"
	"errors"
	"fmt"

	productClient "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/domain"
	repo "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils/response"
)

type cartUseCase struct {
	repo          repo.CartRepository
	productClient productClient.ProductClient
}

func NewCartUseCase(repo repo.CartRepository, productClient productClient.ProductClient) interfaces.CartUseCase {
	return &cartUseCase{
		repo:          repo,
		productClient: productClient,
	}
}

var (
	ErrInvalidProductItemID          = errors.New("invalid product_item_id")
	ErrProductItemAlreadyExistOnCart = errors.New("product_item already exist on cart")
	ErrProductItemOutOfStock         = errors.New("product_item out of stock")
)

func (c *cartUseCase) AddToCart(ctx context.Context, userID, productItemID uint64) error {

	productItem, err := c.productClient.FindProductItemByID(ctx, productItemID)
	if err != nil {
		return err
	}
	if productItem.ID == 0 {
		return ErrInvalidProductItemID
	}

	if productItem.QtyInStock <= 0 {
		return ErrProductItemOutOfStock
	}

	cart, err := c.repo.FindCartByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed find user cart \nerror:%w", err)
	}

	if cart.ID == 0 {
		cart.ID, err = c.repo.SaveCart(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to create cart for user \nerror:%w", err)
		}
	}

	productItemExist, err := c.repo.IsProductItemAlreadyExistInCart(ctx, cart.ID, productItemID)
	if err != nil {
		return fmt.Errorf("failed to check cart_item already exist \nerror:%w", err)
	}
	if productItemExist {
		return ErrProductItemAlreadyExistOnCart
	}

	err = c.repo.SaveCartItem(ctx, cart.ID, productItemID)
	if err != nil {
		return fmt.Errorf("failed to save cart_item \nerror:%w", err)
	}
	return nil
}

func (c *cartUseCase) FindCart(ctx context.Context, userID uint64) (domain.Cart, error) {
	cart, err := c.repo.FindCartByUserID(ctx, userID)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to find cart \nerror:%w", err)
	}
	return cart, nil
}
func (c *cartUseCase) FindCartItem(ctx context.Context, cartID uint64) ([]response.CartItem, error) {

	cartItems, err := c.repo.FindCartItemsByCartID(ctx, cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to find cart_items \nerror:%w", err)
	}

	for i := 0; i < len(cartItems); i++ {
		productItem, err := c.productClient.FindProductItemByID(ctx, cartItems[i].ProductItemID)
		if err != nil {
			return nil, fmt.Errorf("failed to find product_item \nerror:%w", err)
		}
		cartItems[i].ProductName = productItem.Name
		cartItems[i].SKU = productItem.SKU
		cartItems[i].VariationValue = productItem.VariationValue
		cartItems[i].DiscountPrice = productItem.DiscountPrice
		cartItems[i].Price = productItem.Price
		cartItems[i].QtyInStock = productItem.QtyInStock
		cartItems[i].SubTotal = productItem.Price * float64(cartItems[i].Qty)
	}
	return cartItems, nil
}

func (c *cartUseCase) RemoveAllCartItems(ctx context.Context, userID uint64) error {

	cart, err := c.repo.FindCartByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to find cart \nerror:%w", err)
	}

	err = c.repo.RemoveAllCartItems(ctx, cart.ID)

	return err
}

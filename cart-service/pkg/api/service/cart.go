package service

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/usecase"
	usecaseInterface "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartServiceServer struct {
	pb.UnimplementedCartServiceServer
	usecase usecaseInterface.CartUseCase
}

func NewCartServiceServer(usecase usecaseInterface.CartUseCase) pb.CartServiceServer {
	return &CartServiceServer{
		usecase: usecase,
	}
}
func (c *CartServiceServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	utils.LogMessage(utils.Cyan, "AddToCart Invoked")
	err := c.usecase.AddToCart(ctx, req.GetUserId(), req.GetProductId())

	if err != nil {
		utils.LogMessage(utils.Red, "Failed to save cart_item")
		errCode := codes.Internal
		if err == usecase.ErrInvalidProductItemID {
			errCode = codes.InvalidArgument
		}
		return nil, status.Error(errCode, err.Error())
	}

	return &pb.AddToCartResponse{}, nil
}
func (c *CartServiceServer) FindCart(ctx context.Context, req *pb.FindCartRequest) (*pb.FindCartResponse, error) {

	cart, err := c.usecase.FindCart(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	cartItems, err := c.usecase.FindCartItem(ctx, cart.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	outputCartItems := make([]*pb.FindCartResponse_CartItem, len(cartItems))
	var totalPrice float64
	for i := 0; i < len(cartItems); i++ {
		outputCartItems[i] = &pb.FindCartResponse_CartItem{
			ProductItemId:  cartItems[i].ProductItemID,
			ProductName:    cartItems[i].ProductName,
			Sku:            cartItems[i].SKU,
			VariationValue: cartItems[i].VariationValue,
			Price:          cartItems[i].Price,
			DiscountPrice:  cartItems[i].DiscountPrice,
			QtyInStock:     cartItems[i].QtyInStock,
			Qty:            cartItems[i].Qty,
			SubTotal:       cartItems[i].SubTotal,
		}
		totalPrice += cartItems[i].SubTotal
	}
	return &pb.FindCartResponse{
		TotalPrice: totalPrice,
		CartItems:  outputCartItems,
	}, nil
}

func (c *CartServiceServer) RemoveAllCartItems(ctx context.Context, req *pb.RemoveAllCartItemsRequest) (*pb.RemoveAllCartItemsResponse, error) {

	err := c.usecase.RemoveAllCartItems(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RemoveAllCartItemsResponse{}, nil
}

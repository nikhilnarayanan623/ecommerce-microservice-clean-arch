package service

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase"
	usecaseInterface "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderServiceServer struct {
	usecase usecaseInterface.OrderUseCase
	pb.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(usecase usecaseInterface.OrderUseCase) pb.OrderServiceServer {

	return &orderServiceServer{
		usecase: usecase,
	}
}
func (c *orderServiceServer) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {

	shopOrderID, err := c.usecase.PlaceShopOrder(ctx, req.GetUserId())
	if err != nil {
		errCode := codes.Internal
		if err == usecase.ErrCartIsEmpty || err == usecase.ErrCartIsNotValidForOrder {
			errCode = codes.InvalidArgument
		}
		return nil, status.Error(errCode, err.Error())
	}

	return &pb.PlaceOrderResponse{ShopOrderId: shopOrderID}, nil
}

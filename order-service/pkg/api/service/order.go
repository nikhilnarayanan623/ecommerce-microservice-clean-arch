package service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/pb"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase"
	usecaseInterface "github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/request"
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

// Place order for cart items
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

func (c *orderServiceServer) FindAllOrder(ctx context.Context, req *pb.FindAllOrderRequest) (*pb.FindAllOrderResponse, error) {

	pagination := request.Pagination{
		PageNumber: req.GetPageNumber(),
		Count:      req.GetCount(),
	}

	shopOrders, err := c.usecase.FindAllShopOrders(ctx, req.GetUserId(), pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	outputOrders := make([]*pb.FindAllOrderResponse_Orders, len(shopOrders))

	for i, order := range shopOrders {

		orderDate, err := ptypes.TimestampProto(order.OrderDate)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to convert order time.Time to protobuf time stamp")
		}
		outputOrders[i] = &pb.FindAllOrderResponse_Orders{
			ShopOrderId:     order.ID,
			OrderDate:       orderDate,
			OrderTotalPrice: order.OrderTotalPrice,
			Discount:        order.Discount,
		}
	}

	return &pb.FindAllOrderResponse{Orders: outputOrders}, nil
}

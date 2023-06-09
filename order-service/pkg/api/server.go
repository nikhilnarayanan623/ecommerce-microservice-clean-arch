package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs  *grpc.Server
	lis net.Listener
}

func NewServerGRPC(cfg *config.Config, serviceServer pb.OrderServiceServer) (*Server, error) {

	gs := grpc.NewServer()
	pb.RegisterOrderServiceServer(gs, serviceServer)

	lis, err := net.Listen("tcp", cfg.ServiceUrl)
	if err != nil {
		return nil, err
	}

	return &Server{
		gs:  gs,
		lis: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("Order Service Listening.....")
	return c.gs.Serve(c.lis)
}

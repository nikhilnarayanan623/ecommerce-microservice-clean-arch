package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs  *grpc.Server
	lis net.Listener
}

func NewServerGRPC(cfg *config.Config, cartServiceServer pb.CartServiceServer) (*Server, error) {

	gs := grpc.NewServer()
	pb.RegisterCartServiceServer(gs, cartServiceServer)

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
	fmt.Println("Cart Service Listening...")
	return c.gs.Serve(c.lis)
}

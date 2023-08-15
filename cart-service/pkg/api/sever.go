package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs   *grpc.Server
	lis  net.Listener
	port string
}

func NewServerGRPC(cfg *config.Config, cartServiceServer pb.CartServiceServer) (*Server, error) {

	gs := grpc.NewServer()
	pb.RegisterCartServiceServer(gs, cartServiceServer)

	lis, err := net.Listen("tcp", cfg.ServicePort)
	if err != nil {
		return nil, err
	}

	return &Server{
		gs:   gs,
		lis:  lis,
		port: cfg.ServicePort,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("Cart Service Listening on ", c.port)
	return c.gs.Serve(c.lis)
}

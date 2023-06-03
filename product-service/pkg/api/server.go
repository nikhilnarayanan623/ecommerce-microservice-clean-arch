package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs  *grpc.Server
	lis net.Listener
}

func NewServerGRPC(cfg *config.Config, serviceServer pb.ProductServiceServer) (*Server, error) {

	gs := grpc.NewServer()
	pb.RegisterProductServiceServer(gs, serviceServer)

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
	fmt.Println("Product Service Listening.....")
	return c.gs.Serve(c.lis)
}

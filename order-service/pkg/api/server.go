package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs   *grpc.Server
	lis  net.Listener
	port string
}

func NewServerGRPC(cfg *config.Config, serviceServer pb.OrderServiceServer) (*Server, error) {

	gs := grpc.NewServer()
	pb.RegisterOrderServiceServer(gs, serviceServer)

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
	fmt.Println("Order Service Listening On Port ", c.port)
	return c.gs.Serve(c.lis)
}

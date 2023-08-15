package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceServer struct {
	gs   *grpc.Server
	lis  net.Listener
	port string
}

func NewServerGRPC(cfg *config.Config, server pb.AuthServiceServer) (*ServiceServer, error) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", cfg.ServicePort)
	if err != nil {
		return nil, err
	}
	return &ServiceServer{
		gs:   grpcServer,
		lis:  lis,
		port: cfg.ServicePort,
	}, nil
}

func (c *ServiceServer) Start() error {
	fmt.Println("Auth Service Listening on port ", c.port)
	return c.gs.Serve(c.lis)
}

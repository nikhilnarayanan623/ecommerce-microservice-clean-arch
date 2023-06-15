package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceServer struct {
	gs  *grpc.Server
	lis net.Listener
}

func NewServerGRPC(cfg *config.Config, server pb.AuthServiceServer) (*ServiceServer, error) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", cfg.ServiceUrl)
	if err != nil {
		return nil, err
	}
	return &ServiceServer{
		gs:  grpcServer,
		lis: lis,
	}, nil
}

func (c *ServiceServer) Start() error {
	fmt.Println("Auth Service Listening...")
	return c.gs.Serve(c.lis)
}

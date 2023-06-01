package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/user-service/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceServer struct {
	gs  *grpc.Server
	lis net.Listener
}

func NewServerGRPC(cfg *config.Config, server pb.UserServiceServer) (*ServiceServer, error) {

	lis, err := net.Listen("tcp", cfg.ServiceUrl)
	if err != nil {
		return nil, err
	}

	gs := grpc.NewServer()
	pb.RegisterUserServiceServer(gs, server)

	return &ServiceServer{
		gs:  gs,
		lis: lis,
	}, nil
}

func (c *ServiceServer) Start() error {
	fmt.Println("User service listening....")
	return c.gs.Serve(c.lis)
}

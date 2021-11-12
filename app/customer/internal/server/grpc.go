package server

import (
	"context"
	"errors"
	v1 "frame/api/customer/v1"
	"frame/app/customer/internal/conf"
	"frame/app/customer/internal/service"
	"frame/pkg/appmanage"
	"net"

	"google.golang.org/grpc"
)

var errGrpcShutdown = errors.New("error: grpc server has been shutdown")

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func (g *GrpcServer) Serve(ctx context.Context) error {
	select {
	case <-ctx.Done():
		g.server.Stop()
		return errGrpcShutdown
	default:
		return g.server.Serve(g.listener)
	}
}

func NewGrpcServer(service *service.CustomerService, config *conf.GrpcConf) appmanage.GrpcServer {
	server := new(GrpcServer)
	lis, err := net.Listen("tcp", config.Addr())
	server.listener = lis
	if err != nil {
		panic(err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterCustomerServiceServer(grpcServer, service)
	server.server = grpcServer
	return server
}

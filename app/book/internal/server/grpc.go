package server

import (
	"context"
	v1 "frame/api/book/v1"
	"frame/app/book/internal/conf"
	"frame/app/book/internal/service"
	"frame/pkg/appmanage"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func (g *GrpcServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		g.server.Stop()
	}()
	return g.server.Serve(g.listener)
}

func NewGrpcServer(service *service.BookService, config *conf.GrpcConf) appmanage.GrpcServer {
	server := new(GrpcServer)
	lis, err := net.Listen("tcp", config.Addr())
	server.listener = lis
	if err != nil {
		panic(err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterBookServiceServer(grpcServer, service)
	server.server = grpcServer
	return server
}

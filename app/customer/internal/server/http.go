package server

import (
	"context"
	v1 "frame/api/customer/v1"
	"frame/app/customer/internal/conf"
	"frame/app/customer/internal/service"
	"frame/pkg/appmanage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

func (s *HttpServer) Serve(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return s.server.Shutdown(ctx)
	default:
		return s.server.ListenAndServe()
	}
}

func NewHttpServer(service *service.CustomerService, config *conf.HttpConf) appmanage.HttpServer {
	server := new(HttpServer)
	engine := gin.Default()
	v1.RegisterCustomerHttpServer(engine, service)
	server.server = &http.Server{
		Addr:    config.Addr(),
		Handler: engine,
	}
	return server
}

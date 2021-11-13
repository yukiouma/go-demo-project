package server

import (
	"context"
	v1 "frame/api/book/v1"
	"frame/app/book/internal/conf"
	"frame/app/book/internal/service"
	"frame/pkg/appmanage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

func (s *HttpServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.server.Shutdown(ctx)
	}()
	return s.server.ListenAndServe()
}

func NewHttpServer(service *service.BookService, config *conf.HttpConf) appmanage.HttpServer {
	server := new(HttpServer)
	engine := gin.Default()
	v1.RegisterBookHttpServer(engine, service)
	server.server = &http.Server{
		Addr:    config.Addr(),
		Handler: engine,
	}
	return server
}

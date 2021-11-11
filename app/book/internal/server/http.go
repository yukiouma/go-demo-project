package server

import (
	"frame/api/book/v1"
	"frame/app/book/internal/service"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(service *service.BookService) *gin.Engine {
	engine := gin.Default()
	book.RegisterBookHttpServer(engine, service)
	return engine
}

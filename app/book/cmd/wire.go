package main

import (
	"frame/app/book/internal/biz"
	"frame/app/book/internal/data"
	"frame/app/book/internal/server"
	"frame/app/book/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initApp() *gin.Engine {
	panic(wire.Build(
		server.ProvideSet,
		service.ProvideSet,
		data.ProvideSet,
		biz.ProvideSet,
	))
}

// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"frame/app/book/internal/biz"
	"frame/app/book/internal/conf"
	"frame/app/book/internal/data"
	"frame/app/book/internal/server"
	"frame/app/book/internal/service"
	"frame/pkg/appmanage"

	"github.com/google/wire"
)

func initApp(
	db *conf.ConfDB,
	http *conf.HttpConf,
	grpc *conf.GrpcConf,
) *appmanage.AppManage {
	panic(wire.Build(
		server.ProvideSet,
		data.ProvideSet,
		service.ProvideSet,
		biz.ProvideSet,
		appmanage.NewAppManage,
	))
}

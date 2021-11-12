// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"frame/app/customer/internal/biz"
	"frame/app/customer/internal/conf"
	"frame/app/customer/internal/data"
	"frame/app/customer/internal/server"
	"frame/app/customer/internal/service"
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

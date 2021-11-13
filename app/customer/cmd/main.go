package main

import (
	"context"
	"frame/app/customer/internal/conf"
	"frame/pkg/appmanage"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf := conf.GenConf()
	app := initApp(dbConf, httpConf, grpcConf)
	app.Register(&appmanage.RegisterInfo{
		Appid:   "customer:v1",
		AppName: "customer manager service",
	})
	app.Run(ctx, os.Interrupt)
	defer cancel()
}

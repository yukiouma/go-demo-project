package main

import (
	"context"
	"frame/app/book/internal/conf"
	"frame/pkg/appmanage"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf, customerConf := conf.GenConf()
	app := initApp(dbConf, httpConf, grpcConf, customerConf)
	app.Register(&appmanage.RegisterInfo{
		Appid:   "book:v1",
		AppName: "book manager service",
	})
	app.Run(ctx, os.Interrupt)
	defer cancel()
}

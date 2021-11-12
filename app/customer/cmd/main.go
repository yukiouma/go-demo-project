package main

import (
	"context"
	"frame/app/customer/internal/conf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf := conf.GenConf()
	initApp(dbConf, httpConf, grpcConf).Run(ctx)
	defer cancel()
}

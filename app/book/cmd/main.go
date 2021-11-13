package main

import (
	"context"
	"frame/app/book/internal/conf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf, customerConf := conf.GenConf()
	initApp(dbConf, httpConf, grpcConf, customerConf).Run(ctx)
	defer cancel()
}

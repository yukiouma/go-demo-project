package main

import (
	"context"
	"frame/app/book/internal/conf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf := conf.GenConf()
	initApp(dbConf, httpConf, grpcConf).Run(ctx)
	defer cancel()
}

package server

import "github.com/google/wire"

var ProvideSet = wire.NewSet(NewHttpServer, NewGrpcServer)

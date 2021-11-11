# init module frame
new:
	mkdir -p ./api/$(NAME)
	mkdir -p ./app/$(NAME)/cmd
	touch ./app/$(NAME)/cmd/main.go
	touch ./app/$(NAME)/cmd/wire.go
	mkdir -p ./app/$(NAME)/internal
	mkdir -p ./app/$(NAME)/configs
	mkdir -p ./app/$(NAME)/internal/biz
	touch ./app/$(NAME)/internal/biz/biz.go
	mkdir -p ./app/$(NAME)/internal/data
	touch ./app/$(NAME)/internal/data/data.go
	mkdir -p ./app/$(NAME)/internal/service
	touch ./app/$(NAME)/internal/service/service.go
	mkdir -p ./app/$(NAME)/internal/server
	touch ./app/$(NAME)/internal/server/server.go
	mkdir -p ./app/$(NAME)/internal/conf
	touch ./app/$(NAME)/internal/conf/conf.go

# clear all modules
clear:
	rm -rf ./api/*
	rm -rf ./app/*
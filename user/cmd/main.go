package main

import (
	"github.com/MaksKazantsev/grpc_service/user/internal/app"
	"github.com/MaksKazantsev/grpc_service/user/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}

package main

import (
	"learn-microservices-server/src/infra/config"
	"learn-microservices-server/src/server"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conf := config.Make()
	grpcServer := server.NewGRPCServer(

		server.WithConfig(&conf),
	)
	num, _ := strconv.Atoi(conf.Http.Port)
	grpcServer.Run(num)
}

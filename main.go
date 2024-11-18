package main

import (
	"context"
	"log"
	"time"

	usecases "learn-microservices/src/app/usecase"
	bookGrpcUC "learn-microservices/src/app/usecase/books_grpc"
	pickUpUC "learn-microservices/src/app/usecase/pickup"
	"learn-microservices/src/infra/config"
	"learn-microservices/src/infra/persistence/redis"

	"learn-microservices/src/interface/rest"

	booksProto "learn-microservices/src/app/proto/books"
	bookIntegGrpc "learn-microservices/src/infra/integration/books_grpc"

	ms_log "learn-microservices/src/infra/log"

	circuit_breaker_service "learn-microservices/src/infra/circuit_breaker"
	redisService "learn-microservices/src/infra/persistence/redis/service"

	"learn-microservices/src/infra/broker/nats"
	natsPublisher "learn-microservices/src/infra/broker/nats/publisher"

	_ "github.com/joho/godotenv/autoload"
	grpc "google.golang.org/grpc"
)

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	redisClient, err := redis.NewRedisClient(conf.Redis, logger)
	if err != nil {
		panic(err)
	}

	redisSvc := redisService.NewServRedis(redisClient)

	// Connect to the gRPC server
	gRPCConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer gRPCConn.Close()

	// Create a new gRPC client
	gRPCClient := booksProto.NewBookServiceClient(gRPCConn)
	circuitBreaker := circuit_breaker_service.NewCircuitBreakerInstance()
	bookIntegrationGrpc := bookIntegGrpc.NewIntegOpenLibrary(circuitBreaker, gRPCClient)

	Nats := nats.NewNats(conf.Nats, logger)
	publisher := natsPublisher.NewPushWorker(Nats)
	// HTTP Handler
	// the server already implements a graceful shutdown.

	allUC := usecases.AllUseCases{
		PickUpUC:   pickUpUC.NewPickUpUseCase(publisher),
		BookGrpcUC: bookGrpcUC.NewBooksGRPCUseCase(bookIntegrationGrpc, redisSvc),
	}

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		allUC,
		conf.RPS,
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}

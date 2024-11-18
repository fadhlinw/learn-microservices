package server

import (
	"database/sql"
	"learn-microservices-server/db" // Import your db package
	"learn-microservices-server/src/infra/config"
	"learn-microservices-server/src/infra/persistence/redis"
	handler "learn-microservices-server/src/interface/grpc/handlers"

	bookProto "learn-microservices-server/src/app/proto/books"
	usecases "learn-microservices-server/src/app/usecase"
	bookUC "learn-microservices-server/src/app/usecase/books"
	pickUpUC "learn-microservices-server/src/app/usecase/pickup"
	"learn-microservices-server/src/infra/broker/nats"
	pickUpNats "learn-microservices-server/src/infra/broker/nats/consumer/pickup"
	natsPublisher "learn-microservices-server/src/infra/broker/nats/publisher"
	bookInteg "learn-microservices-server/src/infra/integration/books"
	ms_log "learn-microservices-server/src/infra/log"
	redisService "learn-microservices-server/src/infra/persistence/redis/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config *config.Config
	db     *sql.DB // Add db as a field in the server struct
}

type ServerGrpcOption func(*Server)

// NewGRPCServer is constructor
func NewGRPCServer(options ...ServerGrpcOption) *Server {
	server := &Server{}

	for _, option := range options {
		option(server)
	}

	return server
}

func (s *Server) Run(port int) error {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)

	m := make(map[string]interface{})
	m["env"] = s.config.App.Environment
	m["service"] = s.config.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(s.config.Log.Name),
		ms_log.IsProduction(false),
		ms_log.LogAdditionalFields(m))

	// Initialize PostgreSQL connection
	// Use the db package to connect to the database
	db, err := db.ConnectDB() 
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to connect to database")
	}
	// Set the db connection in the server struct
	s.db = db 

	// Ensure the database connection is closed when server shuts down
	defer db.Close()

	redisClient, err := redis.NewRedisClient(s.config.Redis, logger)
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to initialize redis client")
	}
	redisSvc := redisService.NewServRedis(redisClient)

	// Pass db to book integration
	bookIntegration := bookInteg.NewIntegOpenLibrary(s.db)

	Nats := nats.NewNats(s.config.Nats, logger)
	publisher := natsPublisher.NewPushWorker(Nats)

	allUC := usecases.AllUseCases{
		BookUC:   bookUC.NewBooksUseCase(bookIntegration, redisSvc),
		PickUpUC: pickUpUC.NewPickUpUseCase(publisher),
	}

	pickUpNats.NewPickUpWorker(Nats, allUC.PickUpUC)

	handlers := handler.NewHandler(s.config, allUC)

	// register from proto
	bookProto.RegisterBookServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}

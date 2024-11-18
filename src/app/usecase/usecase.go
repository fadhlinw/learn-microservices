package usecase

import (
	bookgRPCUC "learn-microservices/src/app/usecase/books_grpc"
	pickUpUC "learn-microservices/src/app/usecase/pickup"
)

type AllUseCases struct {
	PickUpUC   pickUpUC.PickUpUCInterface
	BookGrpcUC bookgRPCUC.BooksGRPCUCInterface
}

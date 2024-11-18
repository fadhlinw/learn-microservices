package usecase

import (
	bookUC "learn-microservices-server/src/app/usecase/books"
	pickUpUC "learn-microservices-server/src/app/usecase/pickup"
)

type AllUseCases struct {
	BookUC   bookUC.BooksUCInterface
	PickUpUC pickUpUC.PickUpUCInterface
}

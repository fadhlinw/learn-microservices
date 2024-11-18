package pickup

import (
	"encoding/json"
	dto "learn-microservices/src/app/dto/pickup"
	natsPublisher "learn-microservices/src/infra/broker/nats/publisher"
	Const "learn-microservices/src/infra/constants"
	"log"
)

type PickUpUCInterface interface {
	Create(req *dto.ReqPickupDTO) error
}

type pickUpUseCase struct {
	Publisher natsPublisher.PublisherInterface
}

func NewPickUpUseCase(publiser natsPublisher.PublisherInterface) *pickUpUseCase {
	return &pickUpUseCase{
		Publisher: publiser,
	}
}

func (uc *pickUpUseCase) Create(req *dto.ReqPickupDTO) error {
	newData, _ := json.Marshal(req)
	err := uc.Publisher.Nats(newData, Const.BOOKS)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

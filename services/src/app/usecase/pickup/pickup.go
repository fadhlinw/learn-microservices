package pickup

import (
	"encoding/json"
	dto "learn-microservices-server/src/app/dto/pickup"
	natsPublisher "learn-microservices-server/src/infra/broker/nats/publisher"
	Const "learn-microservices-server/src/infra/constants"
	"log"
)

type PickUpUCInterface interface {
	Create(req *dto.ReqPickupDTO) error
	AddDataPickUp(data *dto.ReqPickupDTO)
	GetDataPickUp() []*dto.ReqPickupDTO
}

var dataPickup []*dto.ReqPickupDTO

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

func (uc *pickUpUseCase) AddDataPickUp(data *dto.ReqPickupDTO) {

	dataPickup = append(dataPickup, data)
}

func (uc *pickUpUseCase) GetDataPickUp() []*dto.ReqPickupDTO {

	return dataPickup
}

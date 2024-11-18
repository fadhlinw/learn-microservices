package nats

import (
	"learn-microservices/src/infra/config"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Nats struct {
	Status bool
	Conn   *nats.Conn
}

func NewNats(conf config.NatsConf, logger *logrus.Logger) *Nats {
	var Nats = new(Nats)

	if conf.NatsStatus == "1" {
		Nats.Status = true
	}

	if Nats.Status {
		timeout := 30 * time.Second
		var err error
		Nats.Conn, err = nats.Connect(conf.NatsHost, nats.Timeout(timeout))

		if err != nil {
			logger.Printf("error connecting NATS. %s\n", err.Error())
		}
		log.Println("connected to:", conf.NatsHost)
	}

	return Nats
}

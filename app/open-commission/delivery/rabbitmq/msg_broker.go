package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

type OpenCommissionMessageBroker struct {

	ch          *amqp.Channel
}

func NewCommissionMessageBroker(conn *amqp.Connection) OpenCommissionMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return OpenCommissionMessageBroker{
		ch:          ch,
	}
}
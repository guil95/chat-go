package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	connectionURL = "amqp://rabbitmq:rabbitmq@localhost:5672/"
	QueueBotName  = "bot-stock"
)

func Conn() (*amqp.Connection, error) {
	return amqp.Dial(connectionURL)
}

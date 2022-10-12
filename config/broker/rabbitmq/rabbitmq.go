package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	connectionURL     = "amqp://rabbitmq:rabbitmq@localhost:5672/"
	QueueBotName      = "bot-stock"
	QueueMessagesName = "chat-messages"
)

func Conn() (*amqp.Connection, error) {
	return amqp.Dial(connectionURL)
}

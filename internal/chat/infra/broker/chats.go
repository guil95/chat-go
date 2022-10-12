package broker

import (
	"github.com/guil95/chat-go/config/broker/rabbitmq"
	"github.com/guil95/chat-go/internal/chat"
	"github.com/streadway/amqp"
)

type stockBroker struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewChatBotBroker(conn *amqp.Connection) chat.Broker {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		rabbitmq.QueueMessagesName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return &stockBroker{conn, ch, queue}
}

type Payload struct {
	Message string `json:"message"`
	RoomID  string `json:"roomID"`
}

func (s stockBroker) Send(chat *chat.Chat) error {
	err := s.ch.Publish(
		"",           // exchange
		s.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(chat.ToString()),
		})
	if err != nil {
		return err
	}

	return nil
}

func (s stockBroker) Consume(messageReceived chan []byte) error {
	messages, err := s.ch.Consume(
		rabbitmq.QueueMessagesName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for message := range messages {
			messageReceived <- message.Body
		}
	}()

	return nil
}

package broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guil95/chat-go/config/broker/rabbitmq"
	"github.com/guil95/chat-go/internal/bot"
	"github.com/streadway/amqp"
)

type stockBroker struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewStockBroker(conn *amqp.Connection) bot.Broker {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		rabbitmq.QueueBotName,
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

func (s stockBroker) Send(stock *bot.Stock) error {
	if stock.Value == "N/D" {
		return errors.New("stock value invalid")
	}

	err := s.ch.Publish(
		"",           // exchange
		s.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        s.retrievePayload(stock),
		})
	if err != nil {
		return err
	}

	return nil
}

func (s stockBroker) Consume(messageReceived chan []byte) error {
	messages, err := s.ch.Consume(
		rabbitmq.QueueBotName,
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

func (s stockBroker) retrievePayload(stock *bot.Stock) []byte {
	payload := Payload{
		RoomID:  stock.Room,
		Message: fmt.Sprintf("stock-bot: %v stock is $%v", stock.Code, stock.Value),
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return []byte("")
	}

	return payloadJson
}

package broker

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/guil95/chat-go/internal/chat"
	"github.com/guil95/chat-go/internal/chat/usecase"
)

type Consumer struct {
	broker chat.Broker
	wsConn *websocket.Conn
	uc     *usecase.UseCase
}

func NewChatConsumer(broker chat.Broker, wsConn *websocket.Conn, uc *usecase.UseCase) *Consumer {
	return &Consumer{broker, wsConn, uc}
}

func (c *Consumer) Listen() {
	messageReceived := make(chan []byte)

	_ = c.broker.Consume(messageReceived)

	for {
		select {
		case m := <-messageReceived:
			var cm chat.Chat
			_ = json.Unmarshal(m, &cm)

			_ = c.uc.SaveChat(context.Background(), cm)
		}
	}
}
package broker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/guil95/chat-go/internal/bot"
	"time"
)

type Consumer struct {
	broker bot.Broker
	wsConn *websocket.Conn
}

func NewStockConsumer(broker bot.Broker, wsConn *websocket.Conn) *Consumer {
	return &Consumer{broker, wsConn}
}

type chatMessage struct {
	Message string `json:"message"`
}

func (c *Consumer) Listen() {
	messageReceived := make(chan []byte)

	_ = c.broker.Consume(messageReceived)

	for {
		select {
		case m := <-messageReceived:
			var cm chatMessage
			_ = json.Unmarshal(m, &cm)

			c.wsConn.SetWriteDeadline(time.Now().Add(time.Second * 10))

			err := c.wsConn.WriteMessage(1, []byte(cm.Message))
			if err != nil {
				return
			}

			fmt.Println(fmt.Sprintf("message received %v", string(m)))
		}
	}
}

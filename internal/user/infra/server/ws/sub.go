package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/guil95/chat-go/internal/bot"
	"github.com/guil95/chat-go/internal/chat"
	"go.uber.org/zap"
	"log"
	"strings"
	"time"
)

type Subscription struct {
	conn       *connection
	client     bot.Client
	botBroker  bot.Broker
	chatBroker chat.Broker
	room       string
}

func (s Subscription) readConnectionToHub(username string) {
	c := s.conn
	defer func() {
		hub.unregister <- s
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		m := message{msg, s.room, username}
		hub.broadcast <- m

		_ = s.chatBroker.Send(&chat.Chat{
			Message: string(msg),
			User:    username,
		})

		if stockCode, ok := s.isStockCommand(string(msg)); ok {
			fmt.Println(stockCode)
			result, err := s.client.GetStock(strings.ToLower(stockCode), s.room)
			if err != nil {
				botErrMessage := message{[]byte(err.Error()), s.room, "stock-bot"}
				hub.broadcast <- botErrMessage
			}

			_ = s.chatBroker.Send(&chat.Chat{
				Message: fmt.Sprintf("%v stock is $%v", result.Code, result.Value),
				User:    "stock-bot",
			})

			err = s.botBroker.Send(result)
			if err != nil {
				zap.S().Error(err)
			}
		}
	}
}

func (s *Subscription) writeHubToConnection() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (s *Subscription) isStockCommand(message string) (string, bool) {
	message = strings.TrimSpace(message)
	stockCommandPrefix := "/stock="
	if !strings.HasPrefix(message, stockCommandPrefix) {
		return "", false
	}

	stockCode := strings.TrimPrefix(message, stockCommandPrefix)

	return stockCode, true
}

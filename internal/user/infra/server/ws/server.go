package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/guil95/chat-go/internal/bot"
	"github.com/guil95/chat-go/internal/chat"
)

type (
	// Hub manages the set of active connections for each room
	Hub struct {
		rooms      map[string]map[*connection]bool
		broadcast  chan message
		register   chan Subscription
		unregister chan Subscription
	}

	// message defines the basic message structure
	message struct {
		data     []byte
		room     string
		username string
	}
)

var hub = Hub{
	rooms:      make(map[string]map[*connection]bool),
	broadcast:  make(chan message),
	register:   make(chan Subscription),
	unregister: make(chan Subscription),
}

func init() {
	go hub.Start()
}

func ServeWs(
	roomId,
	username string,
	client bot.Client,
	botBroker bot.Broker,
	chatBroker chat.Broker,
	wsConn *websocket.Conn,
) {
	conn := &connection{ws: wsConn, send: make(chan []byte, 256)}
	sub := Subscription{conn, client, botBroker, chatBroker, roomId}
	hub.register <- sub

	go sub.writeHubToConnection()
	go sub.readConnectionToHub(username)
}

func (h *Hub) Start() {
	var formattedMessage string
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}

			h.rooms[s.room][s.conn] = true

		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}

		case m := <-h.broadcast:
			formattedMessage = fmt.Sprintf("%s: %s", m.username, m.data)
			deliverMessagesToConnections(h, formattedMessage, m.room)
		}
	}
}

func deliverMessagesToConnections(h *Hub, message, room string) {
	connections := h.rooms[room]
	for c := range connections {
		select {
		case c.send <- []byte(message):
		default:
			close(c.send)
			delete(connections, c)
			if len(connections) == 0 {
				delete(h.rooms, room)
			}
		}
	}
}

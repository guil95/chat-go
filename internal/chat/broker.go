package chat

type Broker interface {
	Send(chat *Chat) error
	Consume(messageReceived chan []byte) error
}

package stock

type Broker interface {
	Send(stock *Stock) error
	Consume(messageReceived chan []byte) error
}

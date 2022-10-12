package bot

type Client interface {
	GetStock(code, roomID string) (*Stock, error)
}

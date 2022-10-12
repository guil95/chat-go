package client

import (
	"encoding/csv"
	"fmt"
	"github.com/guil95/chat-go/internal/bot"
	"net/http"
	"strings"
)

const urlTemplate = "https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"

type client struct {
}

func NewClientStock() bot.Client {
	return &client{}
}

func (c client) GetStock(code, roomID string) (*bot.Stock, error) {
	response, err := http.Get(fmt.Sprintf(urlTemplate, strings.ToLower(code)))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	records, err := csv.NewReader(response.Body).ReadAll()
	if err != nil {
		return nil, err
	}

	return &bot.Stock{
		Code:  strings.ToUpper(code),
		Value: records[1][3],
		Room:  roomID,
	}, nil
}

package chat

import "encoding/json"

type Chat struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

func (c *Chat) ToString() string {
	marshal, _ := json.Marshal(c)

	return string(marshal)
}

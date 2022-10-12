package chat

import "context"

type ChatRepository interface {
	SaveChat(ctx context.Context, chat Chat) error
}

package usecase

import (
	"context"
	"testing"

	"github.com/guil95/chat-go/internal/chat"
	"github.com/guil95/chat-go/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUseCase(t *testing.T) {
	t.Run("test save chat without error", func(t *testing.T) {
		repoMock := new(mocks.ChatRepository)
		ctx := context.Background()
		chatMessage := chat.Chat{Message: "Oii", User: "guil95"}

		uc := NewChatBotUseCase(repoMock)

		repoMock.On("SaveChat", ctx, chatMessage).Return(nil)

		err := uc.SaveChat(ctx, chatMessage)
		assert.NoError(t, err)
	})
}

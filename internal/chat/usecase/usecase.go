package usecase

import (
	"context"
	"github.com/guil95/chat-go/internal/chat"
)

type UseCase struct {
	repo chat.ChatRepository
}

func NewChatBotUseCase(repo chat.ChatRepository) *UseCase {
	return &UseCase{repo}
}

func (uc UseCase) SaveChat(ctx context.Context, chat chat.Chat) error {
	return uc.repo.SaveChat(ctx, chat)
}

package usecase

import (
	"context"
	"github.com/guil95/chat-go/internal/user/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type UseCase struct {
	repo domain.Repository
}

func NewUseCase(repo domain.Repository) *UseCase {
	return &UseCase{repo}
}

func (uc UseCase) SaveUser(ctx context.Context, user *domain.User) error {
	u, err := uc.repo.FindByUserName(ctx, user.Username)
	if err != nil && err != mongo.ErrNoDocuments {
		return domain.UserError{}
	}

	if u != nil {
		return domain.UserExists{}
	}

	user.EncodePassword()

	err = uc.repo.SaveUser(ctx, user)
	if err != nil {
		return domain.UserError{}
	}

	return nil
}

func (uc UseCase) Login(ctx context.Context, user *domain.User) error {
	userFromDb, err := uc.repo.FindByUserName(ctx, user.Username)
	if err != nil {
		return domain.UserError{}
	}

	err = user.ComparePassword(userFromDb.Password)
	if err != nil {
		return domain.UserUnauthorized{}
	}

	return nil
}

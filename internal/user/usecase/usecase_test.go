package usecase

import (
	"context"
	"errors"
	"github.com/guil95/chat-go/internal/user/domain"
	"github.com/guil95/chat-go/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FindByUserName = "FindByUserName"
	SaveUser       = "SaveUser"
)

func TestUseCase(t *testing.T) {
	user := &domain.User{Username: "Gui", Password: "123abc"}
	ctx := context.Background()
	t.Run("test login with invalid user should return user error", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		repoMock.On(FindByUserName, ctx, user.Username).Return(nil, errors.New("generic error"))

		uc := NewUseCase(repoMock)
		err := uc.Login(ctx, user)

		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.UserError{})
	})

	t.Run("test login with invalid password return user unauthorized", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		repoMock.On(FindByUserName, ctx, user.Username).Return(&domain.User{Username: "Gui", Password: "anotherPass"}, nil)

		uc := NewUseCase(repoMock)
		err := uc.Login(ctx, user)

		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.UserUnauthorized{})
	})

	t.Run("test login with valid user", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		userDB := &domain.User{Username: "Gui", Password: "123abc"}
		userDB.EncodePassword()
		repoMock.On(FindByUserName, ctx, user.Username).Return(userDB, nil)

		uc := NewUseCase(repoMock)
		err := uc.Login(ctx, user)

		assert.NoError(t, err)
	})

	t.Run("create user with same exist username should return error", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		userDB := &domain.User{Username: "Gui", Password: "123abc"}
		userDB.EncodePassword()
		repoMock.On(FindByUserName, ctx, user.Username).Return(userDB, nil)

		uc := NewUseCase(repoMock)
		err := uc.SaveUser(ctx, user)

		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.UserExists{})
	})

	t.Run("create user with error on find should return error", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		repoMock.On(FindByUserName, ctx, user.Username).Return(nil, errors.New("generic error"))

		uc := NewUseCase(repoMock)
		err := uc.SaveUser(ctx, user)

		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.UserError{})
	})

	t.Run("create user with error should return error", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		repoMock.On(FindByUserName, ctx, user.Username).Return(nil, nil)
		repoMock.On(SaveUser, ctx, user).Return(errors.New("generic error"))

		uc := NewUseCase(repoMock)
		err := uc.SaveUser(ctx, user)

		assert.Error(t, err)
		assert.ErrorIs(t, err, domain.UserError{})
	})

	t.Run("create user with error should return error", func(t *testing.T) {
		repoMock := new(mocks.Repository)

		repoMock.On(FindByUserName, ctx, user.Username).Return(nil, nil)
		repoMock.On(SaveUser, ctx, user).Return(nil)

		uc := NewUseCase(repoMock)
		err := uc.SaveUser(ctx, user)

		assert.NoError(t, err)
	})
}

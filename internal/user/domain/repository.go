package domain

import "context"

type Repository interface {
	SaveUser(ctx context.Context, user *User) error
	FindByUserName(ctx context.Context, userName string) (*User, error)
	FindByUserAndPassword(ctx context.Context, userName, password string) (*User, error)
}

package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

const (
	errorUserExistsMessage   = "user exists"
	errorUserMessage         = "error user"
	errorUnauthorizedMessage = "user unauthorized"
)

type UserError struct{}

func (ue UserError) Error() string {
	return errorUserMessage
}

type UserExists struct{}

func (ue UserExists) Error() string {
	return errorUserExistsMessage
}

type UserUnauthorized struct{}

func (un UserUnauthorized) Error() string {
	return errorUnauthorizedMessage
}

func (u *User) EncodePassword() {
	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	u.Password = string(password)
}

func (u *User) ComparePassword(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
}

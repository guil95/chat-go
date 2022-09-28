package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("encode password and compare without error", func(t *testing.T) {
		userWithEncodedPassword := User{
			Username: "Guilherme",
			Password: "123",
		}

		userWithEncodedPassword.EncodePassword()

		userWithoutEncodedPassword := User{
			Username: "Guilherme",
			Password: "123",
		}

		err := userWithoutEncodedPassword.ComparePassword(userWithEncodedPassword.Password)

		assert.NoError(t, err)
	})

	t.Run("encode password and compare different passwords with error", func(t *testing.T) {
		userWithEncodedPassword := User{
			Username: "Guilherme",
			Password: "1234",
		}

		userWithEncodedPassword.EncodePassword()

		userWithoutEncodedPassword := User{
			Username: "Guilherme",
			Password: "123",
		}

		err := userWithoutEncodedPassword.ComparePassword(userWithEncodedPassword.Password)

		assert.Error(t, err)
	})

	t.Run("encode password and return password encoded", func(t *testing.T) {
		u := User{
			Username: "Guilherme",
			Password: "123",
		}

		u.EncodePassword()

		assert.NotEqual(t, "123", u.Password)
	})

	t.Run("verify user errors message", func(t *testing.T) {
		errorGeneric := UserError{}
		errorExists := UserExists{}
		errorUnauthorized := UserUnauthorized{}

		assert.Equal(t, errorUserMessage, errorGeneric.Error())
		assert.Equal(t, errorUserExistsMessage, errorExists.Error())
		assert.Equal(t, errorUnauthorizedMessage, errorUnauthorized.Error())
	})
}

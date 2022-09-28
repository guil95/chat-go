package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/chat-go/services"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		token, _ := c.Cookie("auth_chat_go")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		if !services.NewJWTService().ValidateToken(token) {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}
}

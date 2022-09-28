package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/chat-go/internal/user/domain"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handleError(ctx *gin.Context, err error) {
	switch err.(type) {
	case domain.UserError:
		ctx.Status(http.StatusInternalServerError)
	case domain.UserUnauthorized:
		ctx.Status(http.StatusUnauthorized)
	case domain.UserExists:
		ctx.JSON(http.StatusUnprocessableEntity, Response{Message: "user already exists"})
	default:
		ctx.Status(http.StatusInternalServerError)
	}
}

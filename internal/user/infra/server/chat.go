package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/chat-go/internal/stock"
	"github.com/guil95/chat-go/internal/user/infra/server/ws"
	"github.com/guil95/chat-go/services"
	"net/http"
)

func (server httpServer) chatRoom(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chat.html", gin.H{})
}

func (server httpServer) chatLobby(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chatLobby.html", gin.H{})
}

type room struct {
	RoomID string `uri:"roomID" binding:"required"`
}

func (server httpServer) chatWs(ctx *gin.Context, client stock.Client, broker stock.Broker) {
	var req room

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	token, _ := ctx.Cookie("auth_chat_go")
	name, _ := services.NewJWTService().GetNameFromToken(token)

	ws.ServeWs(
		ctx.Writer,
		ctx.Request,
		req.RoomID,
		name,
		client,
		broker,
	)
}

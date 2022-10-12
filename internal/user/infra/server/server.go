package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/guil95/chat-go/internal/bot"
	"github.com/guil95/chat-go/internal/user/infra/server/middleware"
	"github.com/guil95/chat-go/internal/user/usecase"
	"net/http"
)

type httpServer struct {
	handler *gin.Engine
	client  bot.Client
	broker  bot.Broker
	uc      *usecase.UseCase
	wsConn  *websocket.Conn
}

func NewHTTPServer(handler *gin.Engine, client bot.Client, broker bot.Broker, uc *usecase.UseCase, wsConn *websocket.Conn) *httpServer {
	return &httpServer{handler, client, broker, uc, wsConn}
}

func (server httpServer) Api() {

	groupUsers := server.handler.Group("/chat", middleware.Auth())
	{
		groupUsers.GET("/:roomID", func(ctx *gin.Context) {
			server.chatRoom(ctx)
		})
		groupUsers.GET("/lobby", func(ctx *gin.Context) {
			server.chatLobby(ctx)
		})
	}

	groupLogin := server.handler.Group("/login")
	{
		groupLogin.GET("", func(ctx *gin.Context) {
			server.loginGet(ctx)
		})
		groupLogin.POST("", func(ctx *gin.Context) {
			server.loginPost(ctx)
		})
	}

	groupUser := server.handler.Group("/users")
	{
		groupUser.POST("", func(ctx *gin.Context) {
			server.createUser(ctx)
		})
	}

	groupIndex := server.handler.Group("", middleware.Auth())
	{
		groupIndex.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}
}

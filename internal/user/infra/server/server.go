package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/chat-go/internal/stock"
	"github.com/guil95/chat-go/internal/user/infra/server/middleware"
	"github.com/guil95/chat-go/internal/user/usecase"
)

type httpServer struct {
	handler *gin.Engine
	client  stock.Client
	broker  stock.Broker
	uc      *usecase.UseCase
}

func NewHTTPServer(handler *gin.Engine, client stock.Client, broker stock.Broker, uc *usecase.UseCase) *httpServer {
	return &httpServer{handler, client, broker, uc}
}

func (server httpServer) Api() {

	groupUsers := server.handler.Group("/chat", middleware.Auth())
	{
		groupUsers.GET("/ws/:roomID", func(ctx *gin.Context) {
			server.chatWs(ctx, server.client, server.broker)
		})
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

}
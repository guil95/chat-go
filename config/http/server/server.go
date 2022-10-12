package server

import (
	"github.com/gorilla/websocket"
	"github.com/guil95/chat-go/config/storages/mongo"
	botBroker "github.com/guil95/chat-go/internal/bot/broker"
	chatBroker "github.com/guil95/chat-go/internal/chat/infra/broker"
	chatRepo "github.com/guil95/chat-go/internal/chat/infra/repository"
	chatUseCase "github.com/guil95/chat-go/internal/chat/usecase"
	"github.com/guil95/chat-go/internal/user/infra/client"
	"github.com/guil95/chat-go/internal/user/infra/repository"
	"github.com/guil95/chat-go/internal/user/infra/server"
	"github.com/guil95/chat-go/internal/user/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guil95/chat-go/config/broker/rabbitmq"
	"go.uber.org/zap"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type ConfigServer struct {
	ws *websocket.Conn
}

func (cs ConfigServer) RunHTTPServer(quit chan os.Signal) {
	handler := gin.Default()
	handler.LoadHTMLGlob("public/templates/*.html")

	rabbitConn, err := rabbitmq.Conn()
	if err != nil {
		zap.S().Errorf("Error on server run %v", err)
		<-quit
	}

	mongoConn := mongo.Connect()
	mongoRepo := repository.NewMongoRepository(mongoConn)

	handler.GET("/chat/ws/:roomID", func(ctx *gin.Context) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		cs.ws, err = upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		bot := botBroker.NewStockConsumer(botBroker.NewStockBroker(rabbitConn), cs.ws)
		go bot.Listen()

		chat := chatBroker.NewChatConsumer(
			chatBroker.NewChatBotBroker(rabbitConn),
			cs.ws,
			chatUseCase.NewChatBotUseCase(chatRepo.NewMongoRepository(mongoConn)),
		)
		go chat.Listen()

		server.ChatWs(
			ctx,
			client.NewClientStock(),
			botBroker.NewStockBroker(rabbitConn),
			chatBroker.NewChatBotBroker(rabbitConn),
			cs.ws,
		)
	})

	s := server.NewHTTPServer(
		handler,
		client.NewClientStock(),
		botBroker.NewStockBroker(rabbitConn),
		usecase.NewUseCase(mongoRepo),
		cs.ws,
	)
	s.Api()

	log.Println("Running http server on :8081 port")
	if err := handler.Run(":8081"); err != nil {
		zap.S().Errorf("Error on server run %v", err)
		<-quit
	}
}

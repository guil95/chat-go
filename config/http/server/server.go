package server

import (
	"github.com/guil95/chat-go/config/storages/mongo"
	"github.com/guil95/chat-go/internal/user/infra/broker"
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

func RunHTTPServer(quit chan os.Signal) {
	handler := gin.Default()
	handler.LoadHTMLGlob("public/templates/*.html")

	rabbitConn, err := rabbitmq.Conn()
	if err != nil {
		zap.S().Errorf("Error on server run %v", err)
		<-quit
	}

	mongoRepo := repository.NewMongoRepository(mongo.Connect())

	s := server.NewHTTPServer(handler, client.NewClientStock(), broker.NewStockBroker(rabbitConn), usecase.NewUseCase(mongoRepo))
	s.Api()

	log.Println("Running http server on :8081 port")
	if err := handler.Run(":8081"); err != nil {
		zap.S().Errorf("Error on server run %v", err)
		<-quit
	}
}

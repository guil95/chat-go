package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/guil95/chat-go/config/http/server"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	cs := server.ConfigServer{}
	
	go cs.RunHTTPServer(quit)

	<-quit

	zap.S().Fatal("Error http server")
}

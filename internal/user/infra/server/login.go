package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/chat-go/internal/user/domain"
	"github.com/guil95/chat-go/services"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type UserPayload struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (server httpServer) createUser(ctx *gin.Context) {
	var u UserPayload
	if err := ctx.ShouldBind(&u); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "VALIDATION_ERROR"})
		return
	}

	err := server.uc.SaveUser(ctx, &domain.User{Username: u.Username, Password: u.Password})
	if err != nil {
		handleError(ctx, err)
		return
	}

	setCookie(ctx, u)
	ctx.JSON(http.StatusCreated, bson.M{})
}

func (server httpServer) loginPost(ctx *gin.Context) {
	var u UserPayload
	if err := ctx.ShouldBind(&u); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "VALIDATION_ERROR"})
		return
	}

	err := server.uc.Login(ctx, &domain.User{Username: u.Username, Password: u.Password})
	if err != nil {
		handleError(ctx, err)
		return
	}

	setCookie(ctx, u)
	ctx.JSON(http.StatusOK, bson.M{})
}

func (server httpServer) loginGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func setCookie(ctx *gin.Context, user UserPayload) {
	token, _ := services.NewJWTService().GenerateToken(user.Username)
	ctx.SetCookie("auth_chat_go", token, 10000, "/", "localhost", true, true)
}

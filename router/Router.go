package router

import (
	"github.com/gin-gonic/gin"
	"im/middlewares"
	"im/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", service.Login)

	r.POST("/send/code", service.SendCode)
	r.POST("/register", service.Register)

	auth := r.Group("/u", middlewares.AuthCheck())
	auth.GET("/user/detail", service.UserDetail)
	auth.GET("/websocket/message", service.WebsocketMessage)
	auth.GET("/chat/list", service.ChatList)
	return r
}

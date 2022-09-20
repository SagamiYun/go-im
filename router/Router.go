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

	auth := r.Group("/u", middlewares.AuthCheck())
	auth.GET("/user/detail", service.UserDetail)

	return r
}

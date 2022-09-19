package router

import (
	"github.com/gin-gonic/gin"
	"im/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", service.Login)
	return r
}

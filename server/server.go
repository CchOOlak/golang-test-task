package server

import (
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	r.GET("/test", test)
	r.POST("/message", message)
	r.GET("/message/list", report)

	return r
}

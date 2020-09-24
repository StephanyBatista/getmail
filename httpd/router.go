package httpd

import (
	"getmail/httpd/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPHandlers() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/environment", handler.EnvironmentGet)
	router.GET("/subscriber", handler.SubscriberGet)
	router.POST("/subscriber", handler.SubscriberPost)
	router.POST("/list", handler.ListPost)

	return router
}

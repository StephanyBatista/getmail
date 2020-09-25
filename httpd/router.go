package httpd

import (
	"getmail/httpd/handler"
	"getmail/infra/data"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func useRepository(repo data.IRepository, handler func(c *gin.Context, repo data.IRepository)) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler(context, repo)
	}
}

func RegisterHTTPHandlers(repo data.IRepository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/environment", handler.EnvironmentGet)
	router.GET("/subscriber", useRepository(repo, handler.SubscriberGet))
	router.POST("/subscriber", useRepository(repo, handler.SubscriberPost))
	router.POST("/list", useRepository(repo, handler.ListPost))

	return router
}

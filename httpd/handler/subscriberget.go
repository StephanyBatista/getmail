package handler

import (
	"getmail/domain/subscribers"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

//SubscriberGet get all subscriber
func SubscriberGet(c *gin.Context, repo data.IRepository) {
	var subscribers []subscribers.Subscriber
	if err := repo.Find(&subscribers); err != nil {
		c.JSON(500, NewDataResponseWithServerError())
		return
	}

	c.JSON(200, subscribers)
}

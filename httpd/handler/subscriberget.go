package handler

import (
	"getmail/domain/subscribers"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

//SubscriberGet get all subscriber
func SubscriberGet(c *gin.Context) {
	var subscribers []subscribers.Subscriber
	if err := data.Repository.Find(&subscribers); err != nil {
		c.JSON(500, NewDataResponseWithServerError())
		return
	}

	c.JSON(200, subscribers)
}

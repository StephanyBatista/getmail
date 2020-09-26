package handler

import (
	"getmail/domain/subscribers"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

//SubscriberGet get all subscriber
func SubscriberGet(c *gin.Context, repo data.IRepository) {

	subscribers, err := repo.Find(&[]subscribers.Subscriber{})
	if err != nil {
		c.JSON(ResponseWithError(err))
		return
	}

	c.JSON(200, subscribers)
}

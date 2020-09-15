package handler

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/domain/subscribers"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

type subscriberRequest struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	ListID string `json:"listid"`
}

//SubscriberPost create a new subscriber
func SubscriberPost(c *gin.Context) {
	requestBody := subscriberRequest{}
	c.Bind(&requestBody)

	model, err := subscribers.New(requestBody.Email, requestBody.Name)
	if err != nil {
		c.JSON(400, NewDataResponseWithError(err))
		return
	}

	if len(requestBody.ListID) > 0 {
		var list lists.List
		if err := data.Repository.First(&list, "ID = ?", requestBody.ListID); err != nil {
			c.JSON(400, NewDataResponseWithError(fmt.Errorf("List not found")))
			return
		}
		model.PutOnList(list.Base.ID)
	}

	if err := data.Repository.Create(model); err != nil {
		c.JSON(500, NewDataResponseWithServerError())
		return
	}

	c.JSON(201, NewDataResponse())
}

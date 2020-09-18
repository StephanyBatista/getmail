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

	if subscribeAlreadySaved(c, requestBody) {
		return
	}

	model, hasError := createNewModel(c, requestBody)
	if !hasError {
		if err := data.Repository.Create(model); err != nil {
			c.JSON(500, NewDataResponseWithServerError())
			return
		}

		c.JSON(201, NewDataResponse())
	}
}

func createNewModel(c *gin.Context, requestBody subscriberRequest) (*subscribers.Subscriber, bool) {
	model, err := subscribers.New(requestBody.Email, requestBody.Name)
	if err != nil {
		c.JSON(400, NewDataResponseWithError(err))
		return nil, true
	}

	return model, putSubscriberOnList(c, requestBody, model)
}

func putSubscriberOnList(c *gin.Context, requestBody subscriberRequest, model *subscribers.Subscriber) bool {
	if len(requestBody.ListID) > 0 {
		var list lists.List
		if err := data.Repository.First(&list, "ID = ?", requestBody.ListID); err != nil {
			c.JSON(400, NewDataResponseWithError(fmt.Errorf("List not found")))
			return true
		} else {
			model.PutOnList(list.Base.ID)
		}
	}
	return false
}

func subscribeAlreadySaved(c *gin.Context, requestBody subscriberRequest) bool {
	var subscribedSaved subscribers.Subscriber
	data.Repository.First(&subscribedSaved, "email = ?", requestBody.Email)

	if len(subscribedSaved.Base.ID) > 0 {
		putSubscriberOnList(c, requestBody, &subscribedSaved)
		if err := data.Repository.Save(&subscribedSaved); err != nil {
			return false
		}
		c.JSON(201, NewDataResponse())
		return true
	}

	return false
}

package handler

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/domain/subscribers"
	"getmail/infra/data"
	"net/http"

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

	if err := subscribeHasAlreadyBeenSaved(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, NewDataResponseWithError(err))
		return
	}

	if err := saveNewSubscriber(requestBody); err != nil {
		c.JSON(http.StatusBadRequest, NewDataResponseWithError(err))
		return
	}

	c.JSON(http.StatusCreated, "")
}

func saveNewSubscriber(requestBody subscriberRequest) error {
	model, err := subscribers.New(requestBody.Email, requestBody.Name)
	if err != nil {
		return err
	}

	putSubscriberOnListIfExist(requestBody.ListID, model)
	data.Repository.Create(model)
	return nil
}

func putSubscriberOnListIfExist(listID string, model *subscribers.Subscriber) error {
	if len(listID) > 0 {
		var list lists.List
		data.Repository.First(&list, "ID = ?", listID)
		if len(list.Base.ID) == 0 {
			return fmt.Errorf("List not found")
		}

		model.PutOnList(list.Base.ID)
	}
	return nil
}

func subscribeHasAlreadyBeenSaved(requestBody subscriberRequest) error {
	var subscribedSaved subscribers.Subscriber
	data.Repository.First(&subscribedSaved, "email = ?", requestBody.Email)

	if len(subscribedSaved.Base.ID) > 0 {
		if err := putSubscriberOnListIfExist(requestBody.ListID, &subscribedSaved); err != nil {
			return err
		}
		data.Repository.Save(&subscribedSaved)
	}

	return nil
}

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
func SubscriberPost(c *gin.Context, repo data.IRepository) {
	requestBody := subscriberRequest{}
	c.Bind(&requestBody)

	if err := subscribeHasAlreadyBeenSaved(requestBody, repo); err != nil {
		c.JSON(http.StatusBadRequest, NewDataResponseWithError(err))
		return
	}

	if err := saveNewSubscriber(requestBody, repo); err != nil {
		c.JSON(http.StatusBadRequest, NewDataResponseWithError(err))
		return
	}

	c.JSON(http.StatusCreated, "")
}

func saveNewSubscriber(requestBody subscriberRequest, repo data.IRepository) error {
	model, err := subscribers.New(requestBody.Email, requestBody.Name)
	if err != nil {
		return err
	}

	putSubscriberOnListIfExist(requestBody.ListID, model, repo)
	repo.Create(model)
	return nil
}

func putSubscriberOnListIfExist(listID string, model *subscribers.Subscriber, repo data.IRepository) error {
	if len(listID) > 0 {
		var list lists.List
		repo.First(&list, "ID = ?", listID)
		if len(list.Base.ID) == 0 {
			return fmt.Errorf("List not found")
		}

		model.PutOnList(list.Base.ID)
	}
	return nil
}

func subscribeHasAlreadyBeenSaved(requestBody subscriberRequest, repo data.IRepository) error {
	var subscribedSaved subscribers.Subscriber
	repo.First(&subscribedSaved, "email = ?", requestBody.Email)

	if len(subscribedSaved.Base.ID) > 0 {
		if err := putSubscriberOnListIfExist(requestBody.ListID, &subscribedSaved, repo); err != nil {
			return err
		}
		repo.Save(&subscribedSaved)
	}

	return nil
}

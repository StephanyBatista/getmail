package handler

import (
	"getmail/domain"
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

	if subscribeExist(requestBody.Email, repo) {
		c.JSON(http.StatusConflict, "")
		return
	}

	if err := saveNewSubscriber(requestBody, repo); err != nil {
		c.JSON(ResponseWithError(err))
		return
	}

	c.JSON(http.StatusCreated, "")
}

func saveNewSubscriber(requestBody subscriberRequest, repo data.IRepository) error {
	model, err := subscribers.New(requestBody.Email, requestBody.Name)
	if err != nil {
		return err
	}

	if len(requestBody.ListID) > 0 {
		if listExist(requestBody.ListID, repo) {
			model.PutOnList(requestBody.ListID)
		} else {
			return domain.NewError("List not exist")
		}
	}

	err = repo.Create(model)
	return err
}

func listExist(listID string, repo data.IRepository) bool {

	if list, _ := repo.First(&lists.List{}, "ID = ?", listID); list != nil {
		return true
	}

	return false
}

func subscribeExist(email string, repo data.IRepository) bool {

	if subscribedSaved, _ := repo.First(&subscribers.Subscriber{}, "email = ?", email); subscribedSaved != nil {
		return true
	}

	return false
}

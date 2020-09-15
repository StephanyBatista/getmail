package handler

import (
	"getmail/domain/lists"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Name string `json:"name"`
}

//ListPost create a new list
func ListPost(c *gin.Context) {

	requestBody := listRequest{}
	c.Bind(&requestBody)

	model, err := lists.New(requestBody.Name)
	if err != nil {
		c.JSON(400, NewDataResponseWithError(err))
		return
	}

	if err := data.Repository.Create(model); err != nil {
		c.JSON(500, NewDataResponseWithServerError())
		return
	}

	c.JSON(201, NewDataResponse())
}

package handler

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Name string `json:"name" form:"name"`
}

//ListPost create a new list
func ListPost(c *gin.Context) {

	requestBody := listRequest{}
	c.Bind(&requestBody)

	var list lists.List
	data.Repository.First(&list, "name = ?", requestBody.Name)
	if len(list.Base.ID) > 0 {
		c.JSON(400, NewDataResponseWithError(fmt.Errorf("List already exists")))
	}

	model, err := lists.New(requestBody.Name)
	if err != nil {
		c.JSON(400, NewDataResponseWithError(err))
		return
	}

	if err := data.Repository.Create(model); err != nil {
		c.JSON(500, NewDataResponseWithServerError())
		return
	}

	c.JSON(201, "")
}

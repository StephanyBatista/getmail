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
func ListPost(c *gin.Context, repo data.IRepository) {

	requestBody := listRequest{}
	c.Bind(&requestBody)

	if requestBody.Name == "" {
		c.JSON(400, DataResponse{Error: "The Name is required"})
		return
	}

	var list lists.List
	repo.First(&list, "name = ?", requestBody.Name)
	if len(list.Base.ID) > 0 {
		c.JSON(400, NewDataResponseWithError(fmt.Errorf("List already exists")))
		return
	}

	model, err := lists.New(requestBody.Name)
	if err != nil {
		c.JSON(400, NewDataResponseWithError(err))
		return
	}

	if err := repo.Create(model); err != nil {
		c.JSON(500, NewDataResponseWithServerError())
		return
	}

	c.JSON(201, "")
}

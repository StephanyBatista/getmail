package handler

import (
	"getmail/domain"
	"getmail/domain/lists"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Name string `json:"name"`
}

//ListPost create a new list
func ListPost(c *gin.Context, repo data.IRepository) {

	requestBody := listRequest{}
	c.Bind(&requestBody)

	model, err := lists.New(requestBody.Name)
	if err != nil {
		c.JSON(ResponseWithError(err))
		return
	}

	if listSaved, _ := repo.First(&lists.List{}, "name = ?", requestBody.Name); listSaved != nil {
		c.JSON(ResponseWithError(domain.NewError("List already exists")))
		return
	}

	if err := repo.Create(model); err != nil {
		c.JSON(ResponseWithError(err))
		return
	}

	c.JSON(201, "")
}

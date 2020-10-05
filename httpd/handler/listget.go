package handler

import (
	"getmail/domain/lists"
	"getmail/infra/data"

	"github.com/gin-gonic/gin"
)

//ListGet get all subscriber
func ListGet(c *gin.Context, repo data.IRepository) {

	lists, err := repo.Find(&[]lists.List{})
	if err != nil {
		c.JSON(ResponseWithError(err))
		return
	}

	c.JSON(200, lists)
}

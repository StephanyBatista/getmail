package handler

import (
	"github.com/gin-gonic/gin"
)

//EnvironmentGet
func EnvironmentGet(c *gin.Context) {

	c.JSON(200, "It working")
}

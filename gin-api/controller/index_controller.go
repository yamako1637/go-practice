package controller

import (
	"gin-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	indexMessage := model.GetIndexMessage()
	c.JSON(http.StatusOK, indexMessage)
}

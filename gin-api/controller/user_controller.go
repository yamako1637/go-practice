package controller

import (
	"gin-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	userData := model.GetUser()
	c.JSON(http.StatusOK, userData)
}

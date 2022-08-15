package main

import (
	"gin-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := router.SetRouter(gin.Default())
	app.Run(":8080")
}

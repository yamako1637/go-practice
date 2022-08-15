package main

import (
	"gin-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	server := router.SetRouter(engine)
	server.Run(":8080")
}

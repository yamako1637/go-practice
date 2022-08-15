package router

import (
	"gin-api/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) *gin.Engine {
	engine.GET("/", controller.IndexController)
	engine.GET("/user", controller.UserController)

	return engine
}

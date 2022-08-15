package router

import (
	"gin-api/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(app *gin.Engine) *gin.Engine {
	app.GET("/", controller.IndexController)
	app.GET("/user", controller.UserController)

	return app
}

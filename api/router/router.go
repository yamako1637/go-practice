package router

import (
	"api/controller"
	"net/http"
)

func Router() {
	http.HandleFunc("/api/user", controller.UserController)
}

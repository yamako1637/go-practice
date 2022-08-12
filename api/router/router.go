package router

import (
	"api/controller"
	"log"
	"net/http"
)

func Router() {
	log.Printf("Server Start")
	http.HandleFunc("/", controller.IndexController)
	http.HandleFunc("/user", controller.UserController)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

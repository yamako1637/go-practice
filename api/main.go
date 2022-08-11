package main

import (
	"api/router"
	"log"
	"net/http"
)

func main() {
	router.Router()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

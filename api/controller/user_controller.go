package controller

import (
	"api/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	log.Printf("/user " + r.Method)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		userData := model.GetUser()
		outputJson, err := json.Marshal(&userData)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(outputJson))
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

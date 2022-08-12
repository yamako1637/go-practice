package controller

import (
	"api/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		indexMessage := model.GetIndexMessage()
		outputJson, err := json.Marshal(&indexMessage)
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

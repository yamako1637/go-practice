package controller

import (
	"api/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
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
		fmt.Printf("処理時間: %vms\n", time.Since(timeStart).Milliseconds())
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

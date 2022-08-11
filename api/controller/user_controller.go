package controller

import (
	"fmt"
	"net/http"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(
			`
				{
					"id":1,
					"name":"hoge"
				}
		`))
		return
	}
}

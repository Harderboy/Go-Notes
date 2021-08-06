package main

import (
	"encoding/json"
	"net/http"
	"middleware_demo/middleware"
)

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	http.HandleFunc("/companies", func(rw http.ResponseWriter, r *http.Request) {
		company := Company{
			ID:      123,
			Name:    "google",
			Country: "USA",
		}
		enc := json.NewEncoder(rw)
		enc.Encode(company)
	})
	http.ListenAndServe(":8080", new(middleware.AuthMiddleWare))
}

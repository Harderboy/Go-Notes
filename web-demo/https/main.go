package main

import (
	"encoding/json"
	"net/http"
)

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// go run /usr/local/go/src/crypto/tls/generate_cert.go -host localhost

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
	http.ListenAndServeTLS(":8080","cert.pem","key.pem",nil)
}

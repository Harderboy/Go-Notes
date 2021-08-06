package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Company struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
}

func main() {
	http.HandleFunc("/companies", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			dec := json.NewDecoder(r.Body)
			company := Company{}
			err := dec.Decode(&company)
			if err != nil {
				log.Println(err.Error())
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			enc := json.NewEncoder(rw)
			err =  enc.Encode(company)
			if err!=nil{
				log.Println(err.Error())
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed) 
		}
	})
	http.ListenAndServe(":8080",nil)
}

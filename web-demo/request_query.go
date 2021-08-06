package main

import (
	"log"
	"net/http"
)

func  main()  {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL
		query := url.Query()  // map[string][]string

		id := query["id"] // 返回一个[]string
		log.Println(id)
		name := query.Get("name")// 返回第一个值 string
		log.Println(name)

	})

	server.ListenAndServe()
}



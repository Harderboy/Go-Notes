package main

import (
	"fmt"
	"net/http"
)

func  main()  {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/header", func(writer http.ResponseWriter, request *http.Request) {
		length := request.ContentLength
		body := make([]byte,length)
		request.Body.Read(body)
		fmt.Fprintln(writer,string(body))
	})

	server.ListenAndServe()
}


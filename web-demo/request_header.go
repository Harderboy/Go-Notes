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
		fmt.Fprintln(writer,request.Header) // map[string][]string
		fmt.Fprintln(writer,request.Header["Accept-Encoding"]) // [gzip, deflate, br]
		fmt.Fprintln(writer,request.Header.Get("Accept-Encoding")) // gzip, deflate, br
	})

	server.ListenAndServe()
}


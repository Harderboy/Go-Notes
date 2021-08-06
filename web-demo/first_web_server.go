package main

import "net/http"

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!"))
	})
	http.ListenAndServe("localhost:8080",nil) // DefaultServeMux
}
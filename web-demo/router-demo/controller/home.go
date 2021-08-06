package controller

import "net/http"

func registerHomeRouter()  {
	http.HandleFunc("/home",handleHome)
}

func handleHome(writer http.ResponseWriter,request *http.Request)  {
	writer.Write([]byte("hello home"))
}
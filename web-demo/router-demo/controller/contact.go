package controller

import "net/http"

func registerContactRouter()  {
	http.HandleFunc("/contact",handleContact)
}

func handleContact(writer http.ResponseWriter,request *http.Request)  {
	writer.Write([]byte("hello contact"))
}
package controller

import "net/http"

func registerAboutRouter()  {
	http.HandleFunc("/about",handleAbout)
}

func handleAbout(writer http.ResponseWriter,request *http.Request)  {
	writer.Write([]byte("hello about"))
}
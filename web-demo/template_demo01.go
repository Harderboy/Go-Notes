package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request){
	t,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/temp.html")
	t.Execute(w,"hello world!")
}

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

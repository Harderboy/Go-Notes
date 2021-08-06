package main

import "net/http"

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type aboutHandler struct{}

func (mh *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

func welcome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome!"))
}

func main() {
	mh := myHandler{}
	a := aboutHandler{}
	server := http.Server{
		//Handler: nil,
		Handler: nil,
		Addr:    "localhost:8080",
	}
	http.Handle("/hello", &mh)
	http.Handle("/about", &a)
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Home!"))
	})
	//http.HandleFunc("/welcome", welcome)
	// 类型转换 http.HandlerFunc 是一个函数类型
	http.Handle("/welcome", http.HandlerFunc(welcome))
	server.ListenAndServe()
	//http.ListenAndServe("localhost:8080",nil)
}

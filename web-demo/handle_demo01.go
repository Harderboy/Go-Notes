package main

import "net/http"

type myHandler struct {}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello world!"))
}
func main(){
	mh:=myHandler{}
	server:=http.Server{
		//Handler: nil,
		Handler: &mh,
		Addr: "localhost:8080",
	}
	server.ListenAndServe()
	//http.ListenAndServe("localhost:8080",nil)
}
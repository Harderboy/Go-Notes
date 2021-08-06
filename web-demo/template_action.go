package main

import (
	"html/template"
	//"math/rand"
	"net/http"
	//"time"
)

func process(w http.ResponseWriter, r *http.Request){
	//t1,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl3.html")
	//rand.Seed(time.Now().Unix())
	//t1.Execute(w,rand.Intn(10)>5)

	//t2,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl2.html")
	//name := []string{"zhangsan","lixi","wangwu"}
	//dayOfWeek:=[]string{}
	//t2.Execute(w,dayOfWeek)


	//t2,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/t1.html","Go入门笔记/web-demo/wwwroot/t2.html")
	//t2.Execute(w,"hello")

	t2,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl.html")
	t2.Execute(w,"hello1")

}

func main() {
	server := http.Server{
		Addr:  "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

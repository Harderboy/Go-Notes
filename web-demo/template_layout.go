package main

import (
	"html/template"
	"log"
	"net/http"
)

func main()  {
	server := http.Server{
		Addr:    "localhost:8080",
	}
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		t,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/index.html","Go入门笔记/web-demo/wwwroot/home.html")

		// 指定模版名
		t.ExecuteTemplate(w,"layout","hello home!")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		t,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/index.html","Go入门笔记/web-demo/wwwroot/about.html")

		// 指定模版
		t.ExecuteTemplate(w,"layout","hello about")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
		t,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/index.html")
		// 指定模版
		e:=t.ExecuteTemplate(w,"layout","hello contact")
		log.Println(e)
		// e为nil时不能调用Error方法
		//log.Println(e.Error())

	})
	server.ListenAndServe()
}

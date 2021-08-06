package main

import (
	"html/template"
	"net/http"
)

func main()  {
	//t,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl.html")

	// 以上代码等于下面两行
	//t:=template.New("Go入门笔记/web-demo/wwwroot/tmpl.html")
	//t,err = t.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl.html")

	// ParseGlob方法
	//template.ParseGlob("*.html")


	server := http.Server{
		Addr:    "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request){
	t,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl.html")
	// 不能指定模版，默认为第一个模版
	t.Execute(w,"hello world!")

	ts,_:=template.ParseFiles("Go入门笔记/web-demo/wwwroot/tmpl.html","Go入门笔记/web-demo/wwwroot/tmpl2.html")
	//err.Error()
	// 指定模版
	ts.ExecuteTemplate(w,"tmpl2.html","hello world!")
}
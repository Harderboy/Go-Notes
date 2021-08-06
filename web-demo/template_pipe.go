package main

import (
	"html/template"
	"net/http"
	"time"
)

/*
如何自定义函数
	template.Funcs(funcMap FuncMap) *Template
	type FuncMap map[string]interface{}
		- value 是函数
		- 可以有任意数量的参数
		- 返回单个值的函数或返回一个值+一个错误的函数
	1。创建一个 FuncMap（map 类型）。
		key 是函数名
		value 就是函数
	2。把 FuncMap 附加到模板
（例子）

*/

func fomateDate(t time.Time) string{
	layout:="2006-01-02"
	return t.Format(layout)
}

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request){
		funcMap:=template.FuncMap{"fdate":fomateDate}
		t:=template.New("temp.html").Funcs(funcMap)
		t.ParseFiles("Go入门笔记/web-demo/wwwroot/temp.html")
		t.Execute(w,time.Now())
	})
	server.ListenAndServe()
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

func loadTemplates() *template.Template  {
	// 随便取的模版名字，模版集
	result := template.New("template")

	// Must()可以包裹一个函数（返回的是一个模板的指针 和 一个错误。)，如果错误不为 nil，那么就 panic
	//t,err:=result.ParseGlob("templates/*.html")
	//template.Must(t,err)
	// ParseGlob 使用模式匹配来解析特定的文件
	// 如果没有匹配到模版文件，err不为nil，Must函数就panic
	template.Must(result.ParseGlob("Go入门笔记/web-demo/templates/*.html"))

	return result
}

func main() {
	templates:=loadTemplates()
	server := http.Server{
		Addr:    "localhost:8080",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// /(斜线后才是文件名)
		fileName:= request.URL.Path[1:]
		//fmt.Println(fileName)
		// 通过模板名来寻找模板，如果没找到就返回 nil
		t:=templates.Lookup(fileName)
		if t!=nil{
			err:=t.Execute(writer,nil)
			if err!=nil{
				log.Fatalln(err.Error())
			}
		}else{
			writer.WriteHeader(http.StatusNotFound)
		}

	})
	http.Handle("/css/",http.FileServer(http.Dir("Go入门笔记/web-demo/wwwroot")))
	http.Handle("/img/",http.FileServer(http.Dir("Go入门笔记/web-demo/wwwroot")))
	server.ListenAndServe()
}

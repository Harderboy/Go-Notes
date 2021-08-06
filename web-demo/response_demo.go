package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
ResponseWriter
	从服务器向客户端返回响应需要使用 ResponseWriter
	ResponseWriter 是一个接口，handler 用它来返回响应
	真正支撑 ResponseWriter 的幕后 struct 是非导出的 http.response

WriteHeader 方法
	WriteHeader 方法接收一个整数类型（HTTP 状态码）作为参数，并把它作为 HTTP 响应的状态码返回
	如果该方法没有显式调用，那么在第一次调用 Write 方法前，会隐式的调用 WriteHeader(http.StatusOK)
	所以 WriteHeader 主要用来发送错误类的 HTTP 状态码
	调用完 WriteHeader 方法之后，仍然可以写入到 ResponseWriter，但无法再修改 header 了
*/
func writeExample(writer http.ResponseWriter, request *http.Request)  {
	str:=`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go Web Programming</title>
</head>
<body>
<h1>hello world!</h1>

</body>
</html>`
	writer.Write([]byte(str))
}


func writeHeaderExample(writer http.ResponseWriter, request *http.Request)  {
	writer.WriteHeader(501)
	fmt.Fprintln(writer,"no such service!")
}

func headerExample(writer http.ResponseWriter, request *http.Request)  {
	// 必须在WriteHeader()方法调用之前修改header，顺序颠倒，就无法再修改header了
	writer.Header().Set("location","www.google.com")
	writer.WriteHeader(302)
}

type Post struct {
	User string
	Threads []string
}

func jsonExample(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type","application/json")
	post :=Post{
		User: "htmoato",
		Threads: []string{"first","second","third"},
	}
	json,_:=json.Marshal(post)
	//writer.Write(json) // {"User":"htmoato","Threads":["first","second","third"]}

	//fmt.Fprintln(writer,json)  //[123 34 85 115 101 114 34 58 34 104 116 109 111 97 116 111 34 44 34 84 104 114 101 97
	// 100 115 34 58 91 34 102 105 114 115 116 34 44 34 115 101 99 111 110 100 34 44 34 116 104 105 114 100 34 93 125]

	fmt.Fprintln(writer,string(json))  // {"User":"htmoato","Threads":["first","second","third"]}

}

//func notFoundExample(writer http.ResponseWriter, request *http.Request)  {
//	http.NotFound()
//}

func  main()  {

	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader",writeHeaderExample)
	http.HandleFunc("/header",headerExample)
	http.HandleFunc("/json",jsonExample)
	server.ListenAndServe()
}


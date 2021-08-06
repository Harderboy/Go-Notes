package main

import "net/http"

func main(){
	/*
	http.FileServer

	●func FileServer(root FileSystem) Handler
	返回一个handler,使用基于root的文件系统来响应请求

	type FileSystem interface {
		Open(name string) (File, error)
	}

	●使用时需要用到操作系统的文件系统，所以还需要委托给:
	type Dir string
	func (d Dir) Open(name string) (File, error)
	*/

	// 测试的时候需要将该目录单独拎出来，作为项目根目录才行，不然会报"404 not found！"

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	http.ServeFile(writer,request,"wwwroot"+request.URL.Path)
	//})
	//http.ListenAndServe(":8080",nil) // DefaultServeMux

	// Simple static webserver 该"Go入门笔记/web-demo/wwwroot"下的文件都可访问
	http.ListenAndServe(":8080",http.FileServer(http.Dir("Go入门笔记/web-demo/wwwroot")))
}

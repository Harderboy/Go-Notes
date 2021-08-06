package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
上传文件
	multipart/form-data 最常见的应用场景就是上传文件（例子）
		首先调用 ParseMultipartForm 方法
		从 File 字段获得 FileHeader，调用其 Open 方法来获得文件
		可以使用 ioutil.ReadAll 函数把文件内容读取到 byte 切片里

FormFile 方法
	上传文件还有一个简便方法：FormFile（例子）
	无需调用 ParseMultipartForm 方法
	返回指定 key 对应的第一个 value
	同时返回 File 和 FileHeader，以及错误信息
	如果只上传一个文件，那么这种方式会快一些

注意点：
	不是所有的 POST 请求都来自 Form
	客户端框架（例如 Angular 等）会以不同的方式对 POST 请求编码：
		jQuery 通常使用 application/x-www-form-urlencoded
		Angular 是 application/json
	ParseForm 方法无法处理 application/json
*/

func process(writer http.ResponseWriter, request *http.Request){

	// 方法1
	//request.ParseMultipartForm(1024)
	//file:=request.MultipartForm.File["uploaded"][0]
	//f,err:=file.Open()

	// 方法2 使用 FormFile 方法
	f,_,err:=request.FormFile("uploaded")
	if err==nil{
		data,err:=ioutil.ReadAll(f)
		if err ==nil{
			fmt.Fprintln(writer,string(data))
		}

	}

}

func main()  {
	server:=&http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}
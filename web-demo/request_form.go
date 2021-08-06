package main

import (
	"fmt"
	"net/http"
)

/*
读取 Form 的值 6种方法：
	Form
	PostForm
	FormValue()
	PostFormValue()
	FormFile()
	MultipartReader()



MultipartForm字段

想要使用MultipartForm这个字段的话，首先需要调用ParseMultipartForm这个方法
	●该方法会在必要时调用ParseForm方法
	●参数是需要读取数据的长度
MultipartForm只包含表单的key-value对
	●返回类型是一个struct而不是map.这个struct里有两个map:
		key是string value是[]string
		空的（key是string value是文件）

FormValue和PostFormValue 方法

	针对 application/x-www-form-urlencoded
	FormValue方法会返回Form字段中指定key对应的第一个value
		无需调用ParseForm或ParseMultipartForm
	PostFormValue 方法也一样，但只能读取PostForm

	FormValue和PostFormValue都会调用ParseMultipartForm方法

小陷阱：
	但如果表单的enctype设为multipart/form-data, 那么即使你调用ParseMultipartForm方法，也无法通过FormValue获得想要的值。
*/
func  main()  {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// h5表单代码
	/*
	<form action="http://localhost:8080/process?first_name=nick" method="post" enctype="multipart/form-data">
	<input type="text" name="first_name">
	<input type="text" name="last_name">
	<input type="submit">
	</form>
	*/
	http.HandleFunc("/process", func(writer http.ResponseWriter, request *http.Request) {
		//request.ParseForm()  // 解析form表单
		//fmt.Fprintln(writer,request.Form)  // 包含form表单和URL中的数据 都是key-value形式，有同样的key的话表单的值靠前，url的值靠后
		//fmt.Fprintln(writer,request.PostForm)  // 只含有form表单中的数据 只支持 application/x-www-form-urlencoded 形式的表单

		//request.ParseMultipartForm(1024)  // 解析 multipart/form-data 形式的表单，1024 是字节数
		//fmt.Fprintln(writer,request.MultipartForm)  // 获取MultipartForm字段，支持 multipart/form-data 形式的表单


		//fmt.Fprintln(writer,request.FormValue("first_name"))  // 获取form字段中指定key对应的第一个value，无需ParseForm/ParseMultipartForm方法去解析
		//fmt.Fprintln(writer,request.PostFormValue("first_name"))  // 只能读取PostForm


		// 表单的enctype设为multipart/form-data
		request.ParseMultipartForm(1024)
		fmt.Fprintln(writer,request.FormValue("first_name"))  // nick
		fmt.Fprintln(writer,request.PostFormValue("first_name"))  // liu

	})

	server.ListenAndServe()
}


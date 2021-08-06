package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger  // 几乎任何东西
	Info *log.Logger  // 重要信息
	Warning *log.Logger  // 警告
	Error *log.Logger  // 错误
)

func init(){
	file,err :=os.OpenFile("Go入门笔记/log/error.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err!= nil{
		log.Fatalln("无法打开错误 log 文件：",err)
	}
	// ioutil.Discard 是一个 io.Writer，它的所有成功执行的 Write 操作都不会产生任何实际的效果。
	//var Discard io.Writer = devNull(0)
	// As of Go 1.16, this value is simply io.Discard.
	var Discard io.Writer = io.Discard
	Trace = log.New(ioutil.Discard,"Trace: ",log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,"Info: ",log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,"Warning: ",log.Ldate|log.Ltime|log.Lshortfile)

	Error =log.New(io.MultiWriter(file,os.Stderr),"Error: ",log.Ldate|log.Ltime|log.Lshortfile)
}
func main(){
	Trace.Println("鸡毛蒜皮的小事")
	Info.Println("一些特别的信息")
	Warning.Println("警告：我是你妈")
	Error.Println("出现了故障")
}
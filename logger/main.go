package main

import (
	"io"
	"log"
	"os"
)

// 自定义logger
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开错误log文件:", err)
	}
	Trace = log.New(io.Discard, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	// defer file.Close()  // 需要注释该行语句，否则错误日志不能写入error.log文件中
}

func main() {
	Trace.Println("小事情")
	Info.Println("特别的信息")
	Warning.Println("警告信息")
	Error.Println("出现了故障")
}

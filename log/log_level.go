package main

import (
	"log"
	"os"
)

func init(){
	log.SetPrefix("LH: ")
	f,err := os.OpenFile("Go入门笔记/log/lh.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	if err!=nil{
		log.Fatalln(err)
	}
	log.SetOutput(f)
	log.SetFlags(log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}

func main(){
	log.Println("1234")

	//log.Fatalln("1234") // 其后会调用os.Exit(1) 退出程序
	// os.Exit(1)

	//log.Panicln("1234")
	//log.Panic("1234")
	//log.Panicf("1234")
}

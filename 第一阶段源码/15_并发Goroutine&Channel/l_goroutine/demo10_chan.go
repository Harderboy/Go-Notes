package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine开始执行。。")
		//time.Sleep(3 * time.Second)
		data := <-ch1 //从ch1中读取数据
		fmt.Println("data：",data)
	}()

	time.Sleep(3*time.Second)
	ch1 <- 10
	fmt.Println("main..over...")


	//ch := make(chan int)
	//ch <- 100 //阻塞

	// 注意书写顺序，防止造成死锁
	// channel必须是成对出现的，通道是goroutine之间的连接，所以通道的发送和接收必须处在不同的goroutine中。
	ch2 := make(chan int)
	go func(){
		ch2 <-200
	}()
	//time.Sleep(3*time.Second)
	data := <-ch2
	fmt.Println(data)

	ch3 := make(chan int)
	go func() {
		data:=<-ch3
		fmt.Println(data)
	}()
	ch3 <- 300
}

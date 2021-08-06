package main

import "fmt"

func main() {
	var ch1 chan  bool
	ch1 = make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println("子goroutine中，i：",i)
		}
		//循环结束后，向通道中写数据，表示要结束了。。
		ch1 <- true
		fmt.Println("结束。。")
	}()

	// 执行接收操作的goroutine将等待直到另一个goroutine尝试向该通道进行发送操作为止。没有接收到数据，将一直处于阻塞状态。
	data := <-ch1
	fmt.Println("main...data-->",data)
	fmt.Println("main...over...")
}

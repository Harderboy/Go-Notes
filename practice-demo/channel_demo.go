/*
通道channel
	通道可以在多个goroutine之间安全的传值
	通道可以作变量、函数参数、结构体字段
	创建通道用make函数，并指定其传输数据的类型
		c:=make(chan int)
	
	通道 channel 发送、接受
	使用左箭头 <- 向通道发送值或者从通道接收值
		向通道发送值： c <- 99
		从通道接收值：r :=<-c
	发送操作会等待知道另一个goroutine尝试对该通道接收操作为止
		执行发送的goroutine在等待期间无法执行其他操作
		未在等待通道操作的goroutine可以继续自由的运行
	执行接收操作的goroutine将等待直到另一个goroutine尝试向该通道进行发送操作为止。
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i <5; i++{
		go sleepyGopher(i,c)
	}
	for i := 0; i <5; i++{
		gopherID := <- c
		fmt.Println("gopher ", gopherID, "has finished sleeping")
	}

}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...",id,"snore...")
	c <- id
}

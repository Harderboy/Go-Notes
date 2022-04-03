package main

import (
	"fmt"
	"time"
)

func main() {

	// var c1, c2, c3 chan int
	c1 := make(<-chan int) // 只读
	// c2 := make(chan<- int) // 只写
	c2 := make(chan int,1)  // 设置缓冲区，否则如果不设置缓存区必须发送方和接收方同时准备就绪才不会阻塞
	c3 := make(chan int) // 可读可写
	var i1, i2 int
	// go func() {
	// 	select {
	// 	case i1 = <-c1:
	// 		fmt.Println("received ", i1, " from c1")
	// 	case c2 <- i2:
	// 		fmt.Println("sent ", i2, " to c2")
	// 	case i3, ok := (<-c3): // same as: i3, ok := <-c3
	// 		if ok {
	// 			fmt.Println("received ", i3, " from c3")
	// 		} else {
	// 			fmt.Printf("c3 is closed, i3:%v ok:%v\n", i3, ok)
	// 		}
	// 	// default:
	// 	// 	fmt.Printf("no communication\n")
	// 	}
	// }()

	i2 = 5
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Printf("c3 is closed, i3:%v ok:%v\n", i3, ok)
		}
	default:
		fmt.Printf("no communication\n")
	}

	// close(c3) // close 只能关闭 只写或者可读可写 的通道
	// c3<-5
	time.Sleep(1 * time.Second)
	fmt.Println("over")
}

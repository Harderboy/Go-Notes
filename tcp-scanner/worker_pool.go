package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("10.196.27.231:%d", p) // 传i会有问题？有很多重复端口号，为什么
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s 关闭了\n", address)
			result<-p
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了！！！！\n", address)
		result<-0
		//wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	result :=make(chan int)
	//var wg sync.WaitGroup
	var openports []int
	var closedports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports,result)
	}

	// 另开一个goroutine 防止造成阻塞，影响result通道更新数据
	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()
	for i:=1;i<1024;i++{
		port:=<-result
		if port==0{
			//closedports = append(closedports,port)
			openports=append(openports,port)
		}else{
			closedports = append(closedports,port)
		}
	}

	// 发送者可以通过关闭信道，来通知接收方不会有更多的数据被发送到channel上。
	close(ports)
	close(result)

	sort.Ints(openports)
	sort.Ints(closedports)

	for _,port:=range openports{
		fmt.Printf("%d open!\n",port)
	}
	for _,port:=range closedports{
		fmt.Printf("%d closed!\n",port)
	}
}

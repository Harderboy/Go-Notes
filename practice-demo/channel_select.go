/*
使用select处理多个通道
	等待不同类型的值
	time.After函数，返回一个通道，该通道在指定时间后会接收到一个值（发送该值的 goroutine 是go运行时的一部分）
	select 和swith有点像
		该语句包含的每个case都持有一个通道，用来发送或接收数据
		select会等待直到某个case分支的操作就绪，然后指向该case
注意：即使已经停止的等待goroutine，但只要main函数还没返回，仍在运行的goroutine将会继续占用内存
select语句在任何case都不满足的情况下，即case处于被阻塞状态时将永远等下去
	如果有 default 子句，则执行该语句。
	如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

nil通道
	不过不使用make初始化通道，那么通道的变量的值就是nil（零值）
	对nil通道进行发送或者接收不会引起panic，但会导致永久阻塞
	对nil通道执行close函数，会引起panic
	nil通道的用处
		对于包含select语句的循环，如果不希望每次循环都等待select所涉及的所有通道，
		那么可以现将某些通道设为nil，等到发送值准备就绪之后，再将通道变成一个非nil值并执行发送操作

阻塞和死锁
	当 goroutine 在等待通道的发送或接收时，我们就说它被阻塞了。
	除了goroutine本身占用少量的内存外，被阻塞的goroutine并不消耗任何其它资源。
		goroutine静静的停在那里，等待导致其阻塞的事情来解除阻塞。
	当一个或多个goroutine因为某些永远无法发生的事情被阻塞时，我们称这种情况为死锁。而出现死锁的程序通常会崩溃或挂起。
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i <5; i++{
		go sleepyGopher(i,c)
	}

	fmt.Println(time.Now())
	//fmt.Println(time.Now())
	timeout := time.After(2 * time.Second)
	for i := 0; i <5; i++{
		select {
		// 两秒内从通道接收数据
		case gopherID :=<- c:
			fmt.Println("gopher ", gopherID, "has finished sleeping")
		// 如果等待两秒后，没有接收到值，会执行该case进行放弃
		// 直到使用<-timer.C发送一个值,timeout即timer.C
		case data:=<-timeout:
			fmt.Println("my patience ran out",data)
			return
		}
	}

}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

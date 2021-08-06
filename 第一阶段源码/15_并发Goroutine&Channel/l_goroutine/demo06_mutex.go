package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

//全局变量，表示票
var ticket = 10 //100张票


var mutex sync.Mutex //mutex 使用：先创建锁头，再在需要加锁的地方进行上锁和解锁

var wg sync.WaitGroup //同步等待组对象
func main() {
	/*
	4个goroutine，模拟4个售票口，


	在使用互斥锁的时候，对资源操作完，一定要解锁。否则会出现程序异常，死锁等问题。
	defer语句
	 */

	 wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait() //main goroutine阻塞等待
	fmt.Println("程序结束了。。。")

	//time.Sleep(5*time.Second)
}

func saleTickets(name string){
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for{
		//上锁
		mutex.Lock() //g2
		if ticket > 0{ //ticket 1 g1
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出：",ticket) // 1
			ticket-- // 0
		}else{
			mutex.Unlock() //条件不满足，也要解锁，防止break跳出去之后，数据一直处于被锁定的状态
			fmt.Println(name,"售罄，没有票了。。")
			break
		}
		mutex.Unlock() //解锁
	}
}

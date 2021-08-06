package main

import (
	"fmt"
	"sync"
	"time"
)

/*
为什么需要读锁：
	当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。
	所以说这里的读锁定（RLock）目的其实是告诉写锁定：有很多人正在读取数据，你给我站一边去，等它们读（读解锁）完你再来写（写锁定）。
*/

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup
func main() {
    // 注意:var rwMutex *sync.RWMutex 只是声明了一个指针变量rwMutex但是没有初始化，
	//指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	rwMutex = new(sync.RWMutex)
	wg = new (sync.WaitGroup)

	//wg.Add(2)
	//多个同时读取
	//go readData(1)
	//go readData(2)

	//wg.Add(3)
	//go writeData(1)
	//go readData(2)
	//go writeData(3)

	wg.Wait()
	fmt.Println("main..over...")
}


func writeData(i int){
	defer wg.Done()
	fmt.Println(i,"开始写：write start。。")
	rwMutex.Lock()//写操作上锁
	fmt.Println(i,"正在写：writing。。。。")
	time.Sleep(3*time.Second)
	rwMutex.Unlock()
	fmt.Println(i,"写结束：write over。。")
}

func readData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始读：read start。。")

	rwMutex.RLock() //读操作上锁
	fmt.Println(i,"正在读取数据：reading。。。")
	time.Sleep(3*time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i,"读结束：read over。。。")
}


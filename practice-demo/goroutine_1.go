/*
在go中，独立的任务叫做goroutine
	与其他语言的中的协程、进程、线程都有相似之处，但是和它们并不完全相同
	它的创建效率非常高
	能直截了当的协同多个并发操作

启动goroutine
	只需在调用前加一个go关键字

每次使用go关键字就会产生一个新的goroutine
	受限与计算机处理单元 goroutine并不是真的在同时运行
	分时技术，轮流执行
	各个goroutine的执行顺序无法确定

goroutine的参数
	向goroutine传入参数就跟函数传递参数一样，参数都是按值传递的（传入的是副本）
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// 启动goroutine
	for i := 0; i <5; i++{
		go sleepyGopher(i)  // 分支线路
	}
	time.Sleep(4 * time.Second)  // 主线路
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}

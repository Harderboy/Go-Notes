package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	临界资源：
	 */
	 a := 1
	 go func() {
	 	a = 2
	 	fmt.Println("goroutine中。。",a)
	 }()

	 a = 3
	 time.Sleep(1) // 让主程序保持不退出
	 fmt.Println("main goroutine...",a)
}
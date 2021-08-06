/*
为了防止panic导致程序崩溃，go提供了recover函数
defer 的动作会在函数返回前执行，即使发生了panic
如果defer的函数调用了recover，panic就回停止，程序将继续运行，类似python中的exception
*/


package main

import "fmt"


func main() {
	
	defer func()  {
		// fmt.Println("defer test")
		if e:=recover();e!=nil{
			fmt.Println("defer test")
		}

	}()

	panic("panic test")

}
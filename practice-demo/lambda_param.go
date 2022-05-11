package main

import (
	"fmt"
)

// 函数当参数传递
// 匿名函数

type cb func(int) int

func testCallBack(x int, f cb){
	f(x)
}

func callback(x int) int{
	fmt.Printf("输出%v\n",x)
	return x
}

func main(){
	testCallBack(2,callback)
	testCallBack(3,func(x int) int{
		fmt.Printf("输出%v\n",x)
		return x
	})

	func(x int) int{
		fmt.Printf("输出%v\n",x)
		return x
	}(4)
	lambda := func(x int) int{
		fmt.Printf("输出%v\n",x)
		return x
	}
	lambda(5)
}
package main

import (
	"fmt"
)

// 闭包
/*
Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
*/


func add(x1,x2 int) func(int,int)(int,int,int){
	i := 0
	return func(x3,x4 int) (int,int,int){
		i+=1
		// fmt.Println(i) 
		return i,x1+x2,x3+x4
	}
}

func main(){

	addFunc := add(1,2)
	fmt.Println(addFunc(3,4))  // i=1
	fmt.Println(addFunc(5,6))  // i=2
}
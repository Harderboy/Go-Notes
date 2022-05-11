package main

import (
	"fmt"
)

var a [5]int
var b [5][5]string

func main()  {
	// 数组初始化的两种方式

	// 方式1
	balance :=[...]int{1,2,3,4}
	fmt.Println(balance)

	// 方式2 指定下标和value
	// arr := [3]int{0:1,2:3}

	// 空值
	arr := [3]int{}
	arr[2] = 10
	fmt.Println(arr)
	
	// 切片初始化，值为空
	var s []int
	fmt.Printf("%T-%v-%v-%v\n",s,len(s),cap(s),s==nil) // s的类型为[]int，s==nil 结果为true
	
	s1 := []int{1,3,4,5,6}
	fmt.Println(s1,len(s1),cap(s1))
	// 添加元素，使用内置的append方法

	s2 := append(s,1,2,3,4)
	fmt.Println("s2: ",s2,len(s2),cap(s2))

    // make([]T, length, capacity)
	// make([]T, length) 此时 length=capacity
	s3 := make([]int,2) // 
	fmt.Printf("s3:%v,length-%v,cap-%v\n",s3,len(s3),cap(s3))

	s4 := make([]int,2,3)
	fmt.Printf("s4:%v,length-%v,cap-%v\n",s4,len(s4),cap(s4))

	// balance 为数组，需要编程切片的形式
	// 数组转换为切片方法 balance[:]

	slice1 := balance[:]
	fmt.Printf("%T",slice1)
	sum := forRange(balance[:])
	fmt.Println(sum)
	sum2 :=forRange(s1)
	fmt.Println(sum2)

}

func forRange(arr []int) int{
	sum :=0
	for index := range arr{
		sum += arr[index]
	}
	return sum
}
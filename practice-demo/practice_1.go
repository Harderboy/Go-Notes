package main

import (
	"fmt"
	"math/rand"
)


func main() {
	fmt.Println("hello world")  // 换行
	fmt.Print("hello ")  // 不换行
	fmt.Print("world!\n")
	fmt.Printf("hello %v \n","go")  // 格式化输出，占位符

	/*
	Author by 
	我是多行注释
	*/

	// 常量
	const distance=500

	//  变量
	// var speed=200
	// var (
	// 	distance2=222
	// 	speed2=688
	// )
	// var distance3,speed3=200,500
	// const hoursPerDay,minPerHour=24,33

	// 赋值
	// var weight=135.00
	// weight=weight*0.25
	// weight*=0.25

	// 自增
	var age=1
	age = age + 1
	age += 1
	// 注意是 age++
	age++
	fmt.Printf("测试自增age：%d\n", age)

	var grade="B"
	var marks int=90

	// switch 的两种用法

	switch marks {
	case 90:
		grade="A"
	case 80:
		grade="B"
	default:
		grade="D"
	}

	switch  {
	case grade=="A":
		fmt.Printf("优秀！\n")
		fallthrough
	case grade=="B",grade=="C":
		fmt.Printf("良好\n")
		fallthrough
	case grade=="D",grade=="E":
		fmt.Printf("良好2\n")
		break
	default:
		fmt.Printf("默认\n")	
	}
	sum:=sum_test(1,2)
	fmt.Printf("测试函数和为：%d\n",sum)

	// switch 使用短声明来声明变量

	switch num:=rand.Intn(10)-1;num{
	case 0:
		fmt.Printf("找到了是0！\n")
	case 1:
		fmt.Printf("找到了是1！\n")
	default:
		fmt.Printf("没找到！\n")

	}

	var l=5
	if l>0 {
		if l>6{
			fmt.Println("l大于6")
		}else if l>4{
			fmt.Println("l大于4")
		}
	}else{
		fmt.Println("l小于0")
	}

	// rune 类型
	var i rune=898
	fmt.Printf("%v-%c\n",i,i)
	var j rune='a'
	fmt.Printf("%v-%c\n",j,j)
}


// 如何定义函数
func sum_test(a,b int) int {
	return a+b
}
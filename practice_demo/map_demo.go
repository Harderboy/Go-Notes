package main

import "fmt"

/* 声明变量，默认 map 是 nil */
var a map[string]int

func main(){
	
	fmt.Printf("%T-%v\n",a,a)

	// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	// 报错 panic: assignment to entry in nil map

	// a["yesterday"] = 1
	// a["today"] = 2
	// a["tomorrow"] = 3
	// fmt.Printf("a:%v\n",a)

	//以上 var a map[string]int 仅仅是声明
	
	// 初始化方法1:
	// 直接创建 使用 make 函数
	b := make(map[string]int)
	// 获取map地址使用 %p
	fmt.Printf("%T-%v-%v-%p\n",b,b,&b,&b)
	
	//方法2 先声明再创建
	var c map[string]int
	c = make(map[string]int)
	fmt.Println(c)

	// 方法3 复合字面值
	d := map[string]int{
		"lihua":2,
		"zhangsan":3,
	}
	fmt.Println(d)

	b["yesterday"] = 1
	b["today"] = 2
	b["tomorrow"] = 3
	fmt.Printf("b:%v\n",b)
	// map中 if ok 的使用,注意ok前面加；
	if value,ok := b["afterday"]; ok {
		fmt.Println(value)
	}

	// 修改
	b["tomorrow"] = 4
	fmt.Printf("b:%v\n",b)
}
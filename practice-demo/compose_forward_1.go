/*
在面向对象的世界中，对象由更小的对象组合而成
术语：对象组合或组合
go通过结构体实现组合（composition）
go还提供了嵌入（embedding）特性，他可以实现方法的转发（forwarding）

组合是一种更简单、灵活的方式
*/


package main

import "fmt"

// 组合
type report struct{
	sol int
	temperation temperation
	location location
}

type temperation struct{
	high, low celsius
}

type location struct{
	lat, long float64
}

type celsius float64

func (t temperation) average() celsius {
	return (t.high + t.low) / 2
}

// 转发
func (r report) average() celsius {
	return r.temperation.average()
}

func main() {
	bradbury :=location{-4.43423, 1374417}
	t := temperation{high: -1.0, low:-78.0}
	fmt.Println(t.average())
	report := report{
		sol: 15,
		temperation: t,
		location: bradbury,
	}
	// fmt.Println(report.temperation.average())
	// 
	fmt.Println(report.average())

	// %+v 输出 struct 字段
	fmt.Printf("%+v\n",report)
	fmt.Printf("a balmy %v C \n",report.temperation.high)
}
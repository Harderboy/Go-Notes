/*
go还可以通过struct（类型）嵌入来实现方法的转发
转发方法：
	在struct中只给定字段类型，不给定字段名即可
	那么被嵌入结构体类型就可以直接使用嵌入类型包含（拥有）的方法，示例如下：

在struct中可以转发任意类型，比如内置的一些类型：int等
*/

package main

import "fmt"

// 另一种 组合
// type report struct{
// 	sol int
// 	temperation temperation
// 	location location
// }
type sol int

type report struct{
	// sol int
	// 嵌入int类型
	// int // 一般不这么写 可以改成 type sol int
	sol
	temperation
	location
}
// 嵌入类型没有指定字段名，只有字段类型，那么改怎么样调用呢？
// 直接使用类型名作为字段名调用即可，go语言会自动给嵌入类型生成一个跟嵌入类型名一样的字段名

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

// func (r report) average() celsius {
// 	return r.temperation.average()
// }

func main() {
	bradbury :=location{-4.43423, 1374417}
	t := temperation{high: -1.0, low:-78.0}
	report := report{
		sol: 15,
		// int: 15,
		temperation: t,
		location: bradbury,
	}
	// report类型直接使用其内嵌入的temperation类型拥有的average方法
	fmt.Println(report.average())

	// 直接使用类型名作为字段名调用即可
	fmt.Printf("a balmy %v C \n",report.temperation.high)
	// 另外一种方式，通过 转发，直接调该嵌入类型下的字段名，类似方法的调用 
	fmt.Printf("a balmy %v C \n",report.high)
}
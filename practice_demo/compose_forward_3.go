/*
命名冲突
优先级
*/

package main

import "fmt"


type sol int

type report struct{
	sol
	temperation
	location
}


type temperation struct{
	high, low celsius
}

type location struct{
	lat, long float64
}

type celsius float64

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0{
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	return 5
}
// 同名方法中 report的优先级最高
// 转发
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	report := report{sol: 5}

	fmt.Println(report.sol.days(1446))
	// 同名方法中 report的优先级最高
	fmt.Println(report.days(1446))
}
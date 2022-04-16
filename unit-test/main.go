package main

import (
	"fmt"
	"./utilss" // 这中写法有点小问题
	// "unit-test/utils"  // cannot find package "unit-test/utils"
)

func main() {
	new := []int{1, 2, 3}
	var b []int
	b = append(b, new...)
	// fmt.Println(new...)  // cannot use new (variable of type []int) as []interface{} value in argument to fmt.PrintlncompilerIncompatibleAssign
	fmt.Println(b)

	fmt.Println(utils.AddTwoNums(1,2))
	s := "ddd"
	fmt.Printf("content:%s", s+" 额外的内容: aaaa")

}

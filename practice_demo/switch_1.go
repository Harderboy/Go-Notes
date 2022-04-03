package main

import "fmt"

func main() {
	a := 10
	// 无匹配项直接跳过
	switch {
	case a > 10:
		fmt.Println("a>10")
	case a < 10:
		fmt.Println("a<10")
	}
	fmt.Println("over")
}

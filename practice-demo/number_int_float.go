package main

import "fmt"

func main() {
	s := []int{2, 2, 1, 1, 1, 2, 2}
	length := len(s)
	fmt.Printf("%T,%v\n", length, length)  // int,7
	fmt.Printf("%T,%v\n", length/2, length/2)  // int,3  除不尽的情况默认向下取整

}

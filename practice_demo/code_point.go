package main

import "fmt"

func main() {
	s:="a1"
	fmt.Printf("%T\n",s[1])  // int8 byte
	fmt.Println(s[0]+s[1]) // 97+49=146
}
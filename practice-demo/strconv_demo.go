package main

import (
	"fmt"
	"strconv"
)

func main() {

	s2 := "100"
	i2,err := strconv.ParseInt(s2,8,64)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("i2: %T, %v,%#v\n",i2,i2,i2)

	// string to int
	i, err := strconv.Atoi("-42")
	if err!=nil{
		fmt.Println(err)
		return
	}
	// int to string
	s := strconv.Itoa(-42)
	// %#v 原样输出
	fmt.Printf("i: %T, %v,%#v\n",i,i,i)
	fmt.Printf("s: %T, %v, %#v\n",s,s,s)
}

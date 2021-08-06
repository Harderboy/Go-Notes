package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var s,sep string
	// s, sep := "", ""
	// os.Args string slice 第一个参数为命令本身

	// for _, args := range os.Args[1:]{
	// 	s += sep + args
	// 	sep = " "
	// }

	// for i:=1;i<len(os.Args);i++{
	// 	s += sep+os.Args[i]
	// 	sep =" "
	// }
	fmt.Println(strings.Join(os.Args[1:]," "))

}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// %s	直接输出字符串或者字节数组
	fmt.Printf("Your name is: %s", text)
}

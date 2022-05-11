package main

import "fmt"

func main() {
	b := make(map[string]int)
	b["a"] = 97
	b["b"] = 98
	fmt.Println(b)
	if value, ok := b["c"]; ok {
		fmt.Println(value)
	}
	for k := range b {
		fmt.Println(k, ":", b[k])
	}
	// 删除key
	delete(b, "b")
	fmt.Println(b)
}

package main

import "fmt"

func main() {
	s := "sdfasf"
	fmt.Printf("%T\n", s[1]) // int8 byte
	for _, v := range s {
		fmt.Printf("%T %v %c\n", v, v, v) // int32 rune
	}

	jewels := "aA"
	stones := "aAAbbbb"

	counts := 0
	resMap := map[byte]bool{}
	for i := range jewels {
		resMap[jewels[i]] = true
	}

	for j := range stones {
		if resMap[stones[j]] {
			counts++
		}
	}
	fmt.Printf("%T", resMap)
}

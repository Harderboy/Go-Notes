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
	fmt.Printf("%T\n", resMap)

	s2 := "欢迎学习Go的len()函数的使用"
	fmt.Println(count(s2)) //14

	fmt.Println(len(s2)) // 37 utf-8
}

func count(str string) int {
    r := []rune(str)
    // count:=0
    // for _:=range str{
        // count++
    // }
    return len(r)
}
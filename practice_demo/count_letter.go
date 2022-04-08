package main

import "fmt"

func main() {
	s:="sdafahsdkfhadkfhdhuhuteryiubbaddgbka"
	countLetter(s)
}

func countLetter(s string) {
	ans := make(map[byte]int)
	for k := range s {
		ans[s[k]]++
	}
	for key, v := range ans {
		fmt.Printf("字母：%c, 次数: %v\n", key,  v)
	}
}

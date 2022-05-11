package main

import "fmt"

// 测试：函数中传递多个参数，其中某个（几个）参数不使用，看是否会报错
// 结果：在函数中传递的参数中可以存在某些参数不使用的情况
func main() {
	fmt.Println(max(1, 2, 3, 4))
}

func max(a, b, c, d int) int {
	// if a > b {
	// 	return a
	// }
	// return b
	return 1
}

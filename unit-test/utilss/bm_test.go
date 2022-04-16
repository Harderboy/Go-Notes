package utils

import (
	"fmt"
	"testing"
)

// 基准测试 demo
func BenchmarkDemo1(b *testing.B) {
	var n int
	// b.N 是测试框架提供
	for i := 0; i < b.N; i++ {
		n++
	}
}

func BenchmarkDemo2(b *testing.B) {

}

// 基准测试 测试函数AddTwoNums
func BenchmarkAddThreeNums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := AddThreeNums(i, i, i)
		fmt.Sprintf("%d", sum)
	}
}

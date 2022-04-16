package utils

import "testing"

// AddTwoNums的测试用例1
// 测试AddTwoNums函数 结果正确
func TestAddTwoNums(t *testing.T) {
	sum := AddTwoNums(1, 2)
	if sum != 3 {
		t.Errorf("AddTwoNums(1, 2)=%d;want 3", sum)
	}
}

// AddTwoNums的测试用例2
// 测试AddTwoNums函数 结果错误
func TestAddTwoNums2(t *testing.T) {
	sum := AddTwoNums(100, 100)
	if sum != 2000 {
		t.Errorf("AddTwoNums(100, 100)=%d;want 2000", sum)
	}
}

func TestAddUp(t *testing.T) {
	sum := AddUp(100)
    if sum != 5050{
        t.Errorf("AddUp(100)=%d;want 5050",sum)
    }
}

func TestAddUpMore(t *testing.T) {
	sum := AddUpMore(1,2,3,4)
    if sum != 10{
        t.Errorf("AddUpMore(1,2,3,4)=%d;want 10",sum)
    }
}

func TestAddUpMore2(t *testing.T) {
	intSlice:=[]int{1,2,3,4}
	sum := AddUpMore(intSlice...)
    if sum != 10{
        t.Errorf("AddUpMore(intSlice)=%d;want 10",sum)
    }
}

package main

import "fmt"

type number struct{
	value int
	valid bool
}

// 使用strcut复合字面值初始化你想要的数据，如果初始化的时候还要做很多事情，可以考虑写一个构造用的函数
// go没有专门的构造函数，但是以new或者New开头的函数，通常是用来构造数据的。比如errors中的New函数
// 构造函数，注意写法
// 构造一个number类型
func newNumber(v int) number {
	return number{value:v, valid:true}
}


/*
fmt在调用Printf或Println的时候，会对传入的参数对象查询并调用一个Stringer接口。
这个Stringer接口有个String() 方法返回一个描述实例自身的字符串。
任何类型只要定义了String()方法，进行Print输出时，就可以得到定制输出。

如果重写了String 方法，那在调用fmt.Println时就会自动去执行String 方法
类似 python中，可以通过_str_()定义输出对象的信息
*/
// 重写了 String方法
func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	// Sprintf：格式化并返回一个字符串而不带任何输出。
	return fmt.Sprintf("%d", n.value)
}

func main() {
	n := newNumber(42)
	fmt.Println(n)  // 42

	e := number{}
	fmt.Println(e)  // not set
}
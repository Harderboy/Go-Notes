package main

import (
	"fmt"
	"strings"
)

type talker interface{
	talk() string
}

func  shout(t talker)  {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string{
	return "nack nack"
}

type laser int

// 指针接收者
func (l *laser) talk() string {
	return strings.Repeat("pew",int(*l))
}

func main()  {
	// 当接口方法的接收者是struct类型时，该类型的变量传指针或者结构体本身都是可以的
	// struct类型会自动解引用
	shout(martian{})
	shout(&martian{})

	pew:=laser(3)
	// 接收者是指针类型时必须要地址
	shout(&pew)
	// shout(pew)

	// 指针接收者
	var t laser  // 声明
	t = 5 // 初始化
	// 方法的接收者和方法的参数在处理指针方面是很相似的
	// go语言在变量在通过点标记法进行调用的时候，自动使用 & 取得变量的内存地址
	// 所以不用写 (&t).talk() 其等同于 t.talk()
	s1 := t.talk()
	fmt.Println(s1)
	s2 := (&t).talk()
	fmt.Println(s2)
}
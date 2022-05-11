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

func (l *laser) talk() string {
	return strings.Repeat("pew",int(*l))
}

func main()  {
	// 当接口方法的接收者是struct类型时，该类型的变量传指针或者结构体本身都是可以的
	// struct类型会自动解引用，以及访问struct时默认会自动加上&，即martian{}等同于&martian{}
	shout(martian{})
	shout(&martian{})

	pew:=laser(3)
	// 接收者是指针类型时必须要地址
	shout(&pew)
	// shout(pew)
}
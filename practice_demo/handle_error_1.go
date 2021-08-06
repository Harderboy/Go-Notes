/*
go 语言允许函数和方法同时返回多个值
按照惯例，函数在返回错误时，最后边的返回值应该用来表示错误

error 内置类型，用来表示错误

defer关键字
	使用defer关键字，go可以确保所有deferred 的动作可以在函数返回前执行。
	可以defer任意的函数和方法
	它不是专门做错误处理的
	他可以消除必须时刻惦记执行资源释放的负担
*/

package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct{
	w io.Writer
	err error
}

func proverbs(name string) error {
	f,err :=os.Create(name)
	if err !=nil{
		return err
	}
	defer f.Close()

	// 写入 
	_,err = fmt.Fprintln(f,"Errors are value.")

	if err!=nil{
		return err
	}
	_,err = fmt.Fprintln(f,"Handle them gracefully")
	return err
}

func main() {
	err := proverbs("proverbs.txt")
	fmt.Println(err)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
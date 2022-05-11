package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type book struct{
	// Marshal 只会将 struct 中被导出字段进行编码，这里要注意，字段名称小写的话json格式输出则为空
	Name string `json:"name"`
	Title string `json:"title"`
}

// 结构体指针
var bookpointer *book

func main()  {
	book1 := book{ Name:"jobs",Title: "biography"}
	fmt.Println("原样输出:",book1)

	b,err := json.Marshal(book1)
	exitOnErr(err)
	if err!= nil{
		fmt.Println(err)
		return 
	}
	fmt.Println("转化为json:",string(b)) 

	// 修改
	book1.Name="mobs"
	fmt.Println("修改:",book1)

	// go提供了内部指针
	// & 不仅能获取结构体地址，还可以获取结构体中指定字段的内存地址
	// 输出结构体地址
	fmt.Printf("地址: %v-%p\n",&book1,&book1) // 地址: &{mobs biography}-0xc000058020

	//输出内部字段地址 直接加上&符号即可
	fmt.Printf("地址: %v-%p\n",&book1.Name,&book1.Name) // 地址: 0xc000058020-0xc000058020

	bookpointer = &book1
	// struct 默认会解引用 直接写var，不用写成*var
	// 使用结构体指针访问结构体成员，使用 "." 操作符：
	fmt.Println(bookpointer.Name)
	fmt.Println((*bookpointer).Name)

}

func exitOnErr(err error){
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
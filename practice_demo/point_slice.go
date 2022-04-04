// slice

/*
slice指向数组

slice是指向数组的窗口，实际上slice在指向数组元素的时候也使用了指针
每个slice内部都会表示一个包含3个元素的结构，他们分别指向
	数组的指针
	slice的容量
	slice的长度
当slice被直接传递至函数或方法时，slice的内部指针就可以对底层数据(数组)进行修改

指向slice的显式指针的唯一作用就是修改slice本身，slice的长度、容量以及起始偏移量
*/

package main

import "fmt"

func reclassify(palnets *[]string){
	// 注意写法
	*palnets = (*palnets)[0:4]
}

func main() {
	// 换行的话注意末尾要有「逗号」 ，
	//missing ',' before newline in composite literalsyntax
	planets := []string{"a","b","c","d","e"}

	reclassify(&planets)
	fmt.Println(planets)
	planets=planets[0:3]
	fmt.Println(planets)
}
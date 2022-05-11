package main

import (
	"fmt"
	"sort"
)

type person struct{
	age int
}

func (p *person) birthday()  {
	// 判断是否为nil保护方法
	if p == nil{
		return
	}
	// struct 指针自动解引用
	// 使用点语法访问其内对应字段即可
	p.age++
}

// 按大小排序
func sortString(s []string, less func(i,j int) bool){
	if less == nil {
		less = func(i,j int)bool {return s[i]<s[j]}
	}
	sort.Slice(s,less)
}

// 按长度排序
func sortStringLen(s []string, less func(i,j int) bool){
	if less == nil {
		less = func(i,j int)bool {return len(s[i])<len(s[j])}
	}
	sort.Slice(s,less)
}

func mirepoix(ingredient []string)[]string{
	return append(ingredient,"onion","carrot","celery")
}

func main() {
	var nowhere *int
	// 指针不为nil，再去解引用
	if nowhere != nil{
		fmt.Println(*nowhere)
	}

	var nobody *person
	fmt.Println(nobody)
	
	/* 
	值为nil的接收者和值为nil的参数在行为上并没有区别，
	所以go语言即使在接收者为nil的情况下，也会继续调用方法
	*/
	nobody.birthday()

	var fn func(int) int
	// 当变量被声明为函数类型时，它的默认值为nil
	fmt.Println(fn == nil)

	food :=[]string{"onion","carrot","celery"}
	sortString(food,nil)
	fmt.Println(food) // [carrot celery onion]
	sortStringLen(food,nil)
	fmt.Println(food) // [onion carrot celery]


	// 如果slice在声明后没有使用复合字面值或内置的make函数进行初始化，那么它的值就是nil
	// 幸运的是 range、len、apend等内置函数都可以正常处理值为nil的slice
	// 虽然空slice 和值为nil的slice并不相等，但是它们通常可以替换使用

	var soup []string
	fmt.Println(soup == nil)

	for _,ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))
	soup = append(soup, "onion","tomato","potato")
	fmt.Println(soup)

	// 虽然空slice 和值为nil的slice并不相等，但是它们通常可以替换使用
	soup2 := mirepoix(nil)
	fmt.Println(soup2)

	var a person
	// struct 的零值不是nil
	fmt.Println(a) // {0}
	// fmt.Println(a==nil)



	// nil map

	var rice map[string]int
	fmt.Println(rice == nil)

	measurement,ok := rice["onion"]
	if ok{
		fmt.Println(measurement)
	}
	
	for ingre,measurement := range rice{
		fmt.Println(ingre,measurement)
	}

	// 声明为接口类型的变量在未被赋值时，接口的零值为nil
	// 对于一个未被复制的接口变量来说，它的接口类型和值都是nil，并且变量本身也是nil
	// 当接口被赋值后，接口就会在内部指向该变量的类型和值

	var v interface{}
	fmt.Printf("%T-%v-%v\n",v,v,v==nil) //<nil>-<nil>-true
	// 在go中，接口类型的变量只有在类型和值都为nil是才等于nil
	// 即使接口变量的值为nil，只要它的类型不是nil，那么该变量就不等于nil
	var p *int
	v = p
	fmt.Printf("%T-%v-%v\n",v,v,v==nil) //*int-<nil>-false

	// 检验接口变量的内部表示
	fmt.Printf("%#v\n",v) // (*int)(nil) 类型：*int 值：nil

}

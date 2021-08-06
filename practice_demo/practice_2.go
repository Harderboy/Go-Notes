package main

import (
	"fmt"
	"math/rand"
	"math"
)

type person struct{
	age int
}


func main()  {
	// 练习几种输入

	// fmt.Println("Please enter your firstname and lastname")
	// var a1,a2 string
	// fmt.Scan(&a1,&a2)
	// fmt.Println("hello ",a1," and ",a2)
	// var a3,a4 string
	// fmt.Scanln(&a3,&a4)
	// fmt.Println("hello ",a3," and ",a4)
	// var a5,a6 string
	// fmt.Scanln(&a5,&a6)
	// fmt.Println("hello",a5,"and",a6)

	// var a1,a2 string
	// fmt.Scanf("%s %s",&a1,&a2)
	// fmt.Println("hello",a1,"and",a2)

	// 练习几种形式的循环

	var pivot=4
	var random=rand.Intn(10)+1
	fmt.Println("随机数为：",random)

	// 形式1 什么都不写 默认为true
	for { 
		if pivot==random{
			fmt.Println("猜对了")
			break
		}else if pivot>random{
			fmt.Println("猜大了")
			pivot--
		}else{
			fmt.Println("猜小了")
			pivot++
		}
	}
	// 形式2 条件为true


	// for true {
	// 	fmt.Println("无限循环")
	// }

	// 形式3 init；condition;post

	// for i:=0;i<=10;i++{
	// 	var random_num=rand.Intn(10)+1
	// 	fmt.Println("随机数为：",random_num)
	// }

	// sum:=0
	// 可直接省略分号
	
	//形式4 只写 condition 省略两个分号，完整版：for ;sum<10;

	// for sum<10 {
	// 	sum+=1
	// 	fmt.Println("随机数为：",sum)
	// }
	// for ;sum<10;{
	// 	sum+=1
	// 	fmt.Println("随机数为：",sum)
	// }

	// a := rand.Int()
	// b := rand.Intn(100)  //生成0-99之间的随机数
	// fmt.Println(a)
	// fmt.Println(b)
	var a float64
	fmt.Println(a)


	// 格式化输出数字
	// third := 1000.0/3
	third := 1.0/3  // 默认为float64
	fmt.Printf("%v\n",third)
	fmt.Printf("%f\n",third)
	fmt.Printf("%.2f\n",third)
	fmt.Printf("%5.2f\n",third)
	fmt.Printf("%05.2f\n",third)  // 可以用0代替默认的空格（补位）

	// 浮点型数字计算
	celsiue := 21.0
	fmt.Println(celsiue / 5.0 * 9.0) // 37.800000000000004
	fmt.Println(9.0 / 5.0 *celsiue)  // 37.800000000000004
	fmt.Println(celsiue*9.0 / 5.0)  // 37.8

	// 打印十六进制的数 使用%x，也可以指定最小宽度和填充
	// %02x 宽度为2，不足的填充为0
	var red, green, blue uint8 = 0, 141, 213
	fmt.Printf("color: # %02x%02x%02x\n",red,green,blue)
	
	// range的使用 三种

	// 切片
	nums := [] int {1,2,3}
	sum :=0
	
	fmt.Println("---------------")
	// 第一种 index value
	for _,value := range nums{
		sum+=value
	}
	fmt.Println("和为：",sum)
	// 第二种 index
	for index := range nums{
		fmt.Println(index)
	}
	
	// 第三种
	for index,_ := range nums{
		fmt.Println(nums[index])
	}
	fmt.Println("---------------")


	// 数组
	arr :=[5] int {1:5,4:6}
	for i,value := range arr{
		fmt.Println(i,value)
	}

	// range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	strings := "go"
	for i,value := range strings{
		fmt.Println(i,value)
		fmt.Printf("%d-%c\n",i,value)
	}

	// 字符串字面值用“”括起来
	// 字符字面值用‘’括起来
	// code := "A"  // string
	code := 'A'  // rune
	fmt.Println(code)

	// 其他数字以 %c 输出
	fmt.Printf("%c\n",67)

	// ` `输出原始字符串字面值

	fmt.Println(`raw string literal\n`)

	sentence := fmt.Sprintf("I am %v years old!", 25)
	fmt.Println(sentence)
	
	answer := 42
	p := &answer // *int
	fmt.Println(*&answer)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Printf("%T\n",p)


	// struct

	// var nobody person
	// fmt.Println(nobody)  // {0}

	var nobody *person
	fmt.Println(nobody) // <nil>
	nobody.birthday()

	fmt.Println("---------------")
	var soup map[string]int
	fmt.Println(soup==nil)
	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}
	// soup 为空无输出，但不报错
	for ingredient,measuremnet:=range soup{
		fmt.Println(ingredient,measuremnet)
		fmt.Println("看上一行是否为空")
	}

	var v interface{}
	fmt.Printf("%T-%v-%v\n", v, v,v == nil)
	var nu *int
	v = nu
	fmt.Printf("%T-%v-%v\n", v, v,v == nil)

	fmt.Printf("%#v\n",v) // 接口类型变量的内部表示：(*int)(nil)

	var s fmt.Stringer
	fmt.Printf("%v-%v\n",s,s==nil) 

	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	

}

// func (p person) birthday(){
// 	p.age++
// }

func (p *person) birthday(){
	if p==nil{
		return
	}
	p.age++
}
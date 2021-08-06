package main

import "fmt"

type Duck interface{
	Quack()
	DuKGo()
}

type Chicken struct{

}

func (c Chicken) IsChicken() bool{
	fmt.Println("我是小鸡")
	return true
}

func (c Chicken) Quack(){
	fmt.Println("嘎嘎")
}
func (c Chicken) DuKGo(){
	fmt.Println("大摇大摆地走")
}

// 鸭子的方法
func DoDuck(d Duck){
	d.Quack()
	d.DuKGo()
}

func main()  {
	var i interface{} = 1
	fmt.Printf("%T-%v\n",i,i)
	// 如果 interface{}中没有定义任何方法，即为空interface，这个任何类型都能满足它。类似泛型（待了解）
	// 当函数参数为interface{}时，可以给它传递任意类型的参数。

	// 小鸡实现了鸭子的所有方法
	// 所以小鸡也是鸭
	// 可以这么写
	c := Chicken{}
	DoDuck(c)

}
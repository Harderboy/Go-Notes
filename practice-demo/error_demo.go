package main

import (
	"fmt"
	"math"
)

func main() {
	radius :=-3.0
	area,err := cicleArea(radius)
	if err!=nil{
		fmt.Println(err)  // 输出的时候 底层自动调用Error方法，类似实现String方法
		if err,ok := err.(*areaError);ok{
			fmt.Printf("半径是 %.2f\n",err.radius)
		}
		return 
	}
	fmt.Println("圆的面积为：", area)

}

// 1.定义一个结构体，表示错误的类型
type areaError struct {
	msg    string
	radius float64
}

//2.实现error接口，就是实现Error()方法
func (a *areaError) Error() string {
	return fmt.Sprintf("error：半径为 %.2f %s", a.radius, a.msg)
}

func cicleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径非法", radius}
		// return 0,areaError{"半径非法",radius}  //错误
	}
	return radius * radius * math.Pi, nil
}

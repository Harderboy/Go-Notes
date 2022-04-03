package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	// 读取channel数据

	// 方式1 range
	for value := range ch {
		fmt.Println("value:", value)
	}

	// 方式2 v,ok:=<-ch
	// for {
	// 	if v, ok := <-ch; ok {
	// 		fmt.Println("v:", v, "ok:", ok)
	// 	} else {
	// 		fmt.Println("break:","v:", v, "ok:", ok)
	// 		break
	// 	}
	// }

	// 同方式2
	// for {
	// 	v2 := <-ch
	// 	if v2 == 3 {
	// 		break
	// 	}
	// 	fmt.Println("v2:", v2)
	// }
}

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc-demo/param"
	"time"
)

// 客户端逻辑实现
func main() {
	//
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	var req float32
	req = 3

	var resp float32

	// 同步调用方式
	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp)
	// 错误调用
	err = client.Call("MathUtil.ddd", req, &resp)
	if err != nil {
		log.Println(err.Error())
	}

	// 多个参数
	myParam := param.AddParam{1, 5}
	var res float32
	err = client.Call("MathUtil.Add", myParam, &res)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Add:", res)

	// 异步调用
	var resp2 float32
	call := client.Go("MathUtil.CalculateCircleArea", req, &resp2, nil)
	// fmt.Println(resp2)
	// reply := <-call.Done
	// fmt.Println(reply)
	// fmt.Println(resp2)
	for {
		select {
		case reply := <-call.Done:
			if call.Error != nil {
				log.Println(err.Error())
				return
			}
			fmt.Println(reply)
			fmt.Println("result:", resp2)
			return
		default:
			fmt.Println("wait...")
			time.Sleep(1 * time.Second)
		}
	}
}

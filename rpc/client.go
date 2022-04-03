package main

import (
	"fmt"
	"net/rpc"
)

// 客户端逻辑实现
func main() {
	//
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	var res float32
	res = 3

	var resp float32

	// 同步调用方式
	err = client.Call("MathUtil.CalculateCircleArea", res, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp)
}

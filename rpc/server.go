package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
	"rpc-demo/param"
)

// 数学计算
type MathUtil struct {
}

// 该方法向外暴露 提供计算圆形面积的服务
func (mu *MathUtil) CalculateCircleArea(req float32, res *float32) error {
	*res = math.Pi * req * req
	return nil
}

func (mu *MathUtil) Add(param param.AddParam, res *float32) error {
	*res = param.Arg1 + param.Arg2
	return nil
}

func main() {
	//1. 初始化指针数据类型
	mathUtil := new(MathUtil)

	//2. 调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}

	//3. 通过该函数把mathUtil提供的服务注册到http协议上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	//4. 在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}

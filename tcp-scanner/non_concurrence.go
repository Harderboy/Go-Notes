package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	start:=time.Now()
	for i:=21;i< 120;i++{
		address:=fmt.Sprintf("10.196.27.231:%d",i)
		conn,err :=net.Dial("tcp",address)
		if err!=nil{
			fmt.Printf("%s 关闭了\n",address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了！！！！\n",address)
	}
	slapsed := time.Since(start)/ 1e9
	fmt.Printf("\n\n %d seconds",slapsed)
}

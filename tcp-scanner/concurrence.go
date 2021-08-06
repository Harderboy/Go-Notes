package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main(){
	start:=time.Now()
	var wg sync.WaitGroup
	for i:=1;i< 65535;i++{
		wg.Add(1)
		go func(j int){
			defer wg.Done()
			address:=fmt.Sprintf("10.196.27.231:%d",j)//传i会有问题？有很多重复端口号，为什么
			conn,err :=net.Dial("tcp",address)
			if err!=nil{
				fmt.Printf("%s 关闭了\n",address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了！！！！\n",address)
		}(i)
	}
	wg.Wait()
	slapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds",slapsed)
}

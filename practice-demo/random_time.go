package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())

	for i :=0;i<10;i++{
		//生成10个0-99之间的随机数
		year := rand.Intn(100)+1+2000 // 2001~2100年间的年份
		month := rand.Intn(12)+1
		daysInMonth := 31 
	
		switch month {
		case 2:
			// 闰年判断条件：
			// 1.年份能被4整除，但不能被100整除；
			// 2.能被400整除
			if (year % 4 == 0 && year % 100 !=0) || year % 400==0{
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 1,3,5,7,8,10,12:
			daysInMonth = 31
		case 4,6,9,11:
			daysInMonth = 30
		}
		fmt.Println(era,year,month,daysInMonth)
	}

}

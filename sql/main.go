package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go_demo"
)

func main() {
	connStr := "root:123456@tcp(127.0.0.1:3306)/go_demo"
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()

	//err = db.Ping()
	//if err != nil{
	//	log.Fatalln(err)
	//}

	fmt.Println("Connected!")

	log.Println(db ==nil)
	// 查询一笔
	// 和以上db保持一致，注意getOne方法引入的是全局变量
	one,err:=getOne(1)
	if err!=nil{
		log.Println(err.Error())
	}
	fmt.Println(one)
	two,err:=getOne(2)
	if err!=nil{
		log.Println(err.Error())
	}
	fmt.Println(two)
	//fmt.Printf("%v-%#v\n",two,two)

	// 查询多笔
	apps,err:=getMany(0)
	if err!=nil{
		log.Fatalln(err.Error())
	}
	fmt.Println(apps)

	// 更新操作
	a,err:=getOne(1)
	fmt.Println(a)
	if err!=nil{
		log.Println(err.Error())
	}
	a.name+="455"
	a.order++

	err=a.Update()
	if err!=nil{
		log.Println(err.Error())
	}
	a1,err:=getOne(1)
	fmt.Println(a1)

	a2:=app{
		name:"potato",
		status: 1,
		order: 3,
		level:6,
	}
	err=a2.Insert()
	if err!=nil{
		log.Println(err.Error())
	}
}

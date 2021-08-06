package main

import (
	"fmt"
	"log"
)

func getOne(id int )(a app,err error)  {
	a=app{}
	//log.Println(db ==nil)
	// order是关键字，使用反引号进行转义
	err=db.QueryRow("select `id`,`name`,`status`,`level`,`order` from `go_demo` where `id`=?",id).Scan(&a.ID,
		&a.name,&a.status,&a.level,&a.order)
	// 上面已经给a和err赋值了，这里的返回不用再显示写出了
	// return a,err
	return
}


func getMany(id int)(apps []app,err error)  {
	rows,err :=db.Query("select `id`,`name`,`status`,`level`,`order` from `go_demo` where `id`>?",id)
	for rows.Next(){
		a:=app{}
		err=rows.Scan(&a.ID,&a.name,&a.status,&a.level,&a.order)
		if err!=nil{
			log.Fatalln(err.Error())
		}
		apps=append(apps,a)
	}
	return
}

func (a *app) Update() (err error){
	_,err=db.Exec("update `go_demo` set name=?,`order`=? where id=?",a.name,a.order,a.ID)
	if err!=nil{
		fmt.Println(err.Error())
	}
	return
}

func (a *app) Insert() (err error){
	sqlStr :="insert into `go_demo` (`name`,`status`,`level`,`order`)values(?,?,?,?);select ifNull(SCOPE_IDENTITY(),-1)"
	stmt,err :=db.Prepare(sqlStr)
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer stmt.Close()

	//_,err=stmt.Exec(a.name,a.status,a.level,a.order)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//return

	// sql server语法
	// Prepare方法创建一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
	// 体会 stmt *Stmt 的用法
	//查询最后插入数据的id字段： select isNull(SCOPE_IDENTITY(),-1)
	//sql2="insert into `go_demo` (`name`,`status`,`level`,`order`)values(?,?,?,?);select isNull(SCOPE_IDENTITY(),-1);"
	//stmt,err :=db.Prepare(sql2)
	//err=stmt.QueryRow(a.name,a.status,a.level,a.order).Scan(&a.ID)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//return
}

func (a *app) Delete() (err error){
	_,err=db.Exec("delete from `go_demo` where id=?",a.ID)
	if err!=nil{
		fmt.Println(err.Error())
	}
	return
}
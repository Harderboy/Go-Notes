package main

import (
	db "gin-webservice-demo/databases"
	. "gin-webservice-demo/router"
)

func main() {
	//当整个程序完成之后关闭数据库连接
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8080")
}

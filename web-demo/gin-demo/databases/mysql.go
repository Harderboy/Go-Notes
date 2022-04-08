package databases

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//因为我们需要在其他地方使用SqlDB这个变量，所以需要大写代表public
var SqlDB *sql.DB

//初始化方法
func init() {
	var err error
	// 使用sql.Open()方法会创建一个数据库连接池db。这个地步不是数据库连接，它是一个连接池，只有当真正的数据库通信的时候才创建连接。
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/person?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

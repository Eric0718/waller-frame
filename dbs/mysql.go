package dbs

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
)

func InitDB(host, port, database, username, password, charset string) *gorm.DB {
	fmt.Println(host, port, database, username, password, charset, "===aaa===>")
	// 字符串拼接
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	fmt.Println("数据库连接:", sqlStr)

	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("打开数据库失败", err)
		panic("打开数据库失败" + err.Error())
	}
	return db
}

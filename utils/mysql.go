package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// 包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	dsn := "root:root@tcp(localhost:3306)/nvideo"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)  //连接池最大允许的空闲连接数，超过的连接会被连接池关闭。
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数

}

// 获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
// 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return db
}

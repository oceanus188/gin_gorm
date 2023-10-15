package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/curd-list?charset=utf8mb4&parseTime=True&loc=Local"
	Db = new(gorm.DB)
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	//数据库连接池
	sqlDB, _ := Db.DB()
	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	//设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

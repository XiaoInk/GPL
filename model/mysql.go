/*
	Created by XiaoInk at 2021/04/23 18:44:55
	GitHub: https://github.com/XiaoInk
*/

package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	var err error

	db, err = gorm.Open(mysql.Open(Config.MySQL+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "oms_", // 表名前缀
			SingularTable: true,   // 启用单数表名
		},
	})

	if err != nil {
		log.Fatalln("Gorm.Open.err", err.Error())
	}

	// 设置连接池
	sqlDB, _ := db.DB()
	// 最大空闲连接
	sqlDB.SetMaxIdleConns(10)
	// 最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接最大可复用时间
	sqlDB.SetConnMaxLifetime(time.Second * 30)
}

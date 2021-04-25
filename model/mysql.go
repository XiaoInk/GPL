/*
	Created by XiaoInk at 2021/04/23 18:44:55
	GitHub: https://github.com/XiaoInk
*/

package model

import (
	"log"
	"time"

	"github.com/XiaoInk/GPL/model/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Getdb() *gorm.DB {
	return db
}

func init() {
	var err error

	db, err = gorm.Open(mysql.Open(Config.MySQL.URI+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Config.MySQL.TablePrefix, // 表名前缀
			SingularTable: true,                     // 启用单数表名
		},
	})

	if err != nil {
		log.Fatalln("Gorm.Open.err", err.Error())
	}

	// 设置连接池
	sqldb, _ := db.DB()
	// 最大空闲连接
	sqldb.SetMaxIdleConns(10)
	// 最大打开连接数
	sqldb.SetMaxOpenConns(100)
	// 连接最大可复用时间
	sqldb.SetConnMaxLifetime(time.Second * 30)

	// 数据表迁移
	err = db.AutoMigrate(&table.User{})
	if err != nil {
		log.Fatalln("Gorm.AutoMigrate.err", err.Error())
	}

	// 用户表初始化
	user := table.User{}
	user.Init(db)
}

package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建数据库句柄
var GlobalDB *gorm.DB

type People struct {
	gorm.Model
	Name string
	Age  int
}

func InitDB() {
	dns := "pjh:Pjh_921122@tcp(120.76.141.85:3306)/PJH_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		fmt.Printf("gorm err: %s.\n", err)
		// return err
	}
	//连接池设置
	//设置初始化数据库链接个数
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	GlobalDB = db
	//表的创建
	fmt.Printf("gorm AutoMigrate err: %s.\n", db.AutoMigrate(&People{}).Error())

}

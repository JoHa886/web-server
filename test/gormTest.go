package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type People struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dns := "root:20211201@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		fmt.Printf("gorm err: %s.\n", err)
		return
	}
	// 创建数据库表
	db.AutoMigrate(&People{})
	db.Create(&People{Name: "PJH", Age: 28})

}

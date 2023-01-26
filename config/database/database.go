package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/youbride_dev?charset=utf8mb4&parseTime=True&loc=Local")

	// 単数系のテーブルを許可する
	db.SingularTable(true)

	if err != nil {
		panic(err.Error())
	}

	return db
}

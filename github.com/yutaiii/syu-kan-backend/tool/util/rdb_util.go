package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConn() *gorm.DB {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(syu-kan-db:3306)"
	DBNAME := "syu_kan"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

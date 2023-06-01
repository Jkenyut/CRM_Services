package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GormMysql() *gorm.DB {
	var db, err = gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/crm_service?charset=utf8mb4&parseTime=True&loc=UTC"), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}

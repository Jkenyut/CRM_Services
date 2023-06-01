package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func GormMysql() *gorm.DB {
	var db, err = gorm.Open(mysql.Open(os.Getenv("CONNECT_DB")), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}

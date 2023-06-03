package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func GormMysql() *gorm.DB {
	var db, err = gorm.Open(mysql.Open(os.Getenv("CONNECT_DB")), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db = db.Debug()
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}

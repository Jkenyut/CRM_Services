package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormMysql() *gorm.DB {
	dsn := fmt.Sprint("root@tcp(127.0.0.1:3306)/crm_bootcamp?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Check for errors when opening the connection.
	if err := db.Error; err != nil {
		panic("GORM error: " + err.Error())
	}
	db.Debug()
	return db

}

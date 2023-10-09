package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func GormMysql() *gorm.DB {
	dsn := os.Getenv("CONNECT_DB")
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

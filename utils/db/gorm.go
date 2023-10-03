package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GormMysql() *gorm.DB {
	dsn := fmt.Sprint("root@tcp(", os.Getenv("CONNECT_DB"), ")/crm_bootcamp?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

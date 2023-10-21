package connection

import (
	"crm_service/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type InterfaceConnection interface {
	Init()
	GetConnectionDB() *gorm.DB
}

type ManagerConnection struct {
	config   *config.Config
	database *gorm.DB
}

func NewConnection(conf *config.Config) InterfaceConnection {
	return &ManagerConnection{
		config: conf,
	}
}

func (c *ManagerConnection) GetConfig() *config.Config {
	return c.config
}

func (c *ManagerConnection) GetConnectionDB() *gorm.DB {
	return c.database.Debug()
}

func (c *ManagerConnection) Init() {
	dsn := c.GetConfig().Database.URL
	log.Println("dsn", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Print("Failed to connect to the database: " + err.Error())
		panic("Failed to connect to the database: " + err.Error())
	}

	// Check for errors when opening the connection.
	if err = db.Error; err != nil {
		log.Print("connect database error: " + err.Error())
		panic("connect database error: " + err.Error())
	}
	c.database = db
}

//
//func GormMysql() *gorm.DB {

//}

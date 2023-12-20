package connection

import (
	"crm_service/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type InterfaceConnection interface {
	Init(isuat bool)
	GetConnectionDB() *gorm.DB
}

type ManagerConnection struct {
	config   *config.Config
	database *gorm.DB
	isuat    bool
}

func NewConnection(conf *config.Config, isuat bool) InterfaceConnection {
	return &ManagerConnection{
		config: conf,
		isuat:  isuat,
	}
}

func (c *ManagerConnection) GetConnectionDB() *gorm.DB {
	return c.database.Debug()
}

func (c *ManagerConnection) GetConfigConnection() *config.Config {
	return c.config
}

func (c *ManagerConnection) Init(isuat bool) {

	if len(c.config.KeyAES) != 16 && len(c.config.KeyAES) != 24 && len(c.config.KeyAES) != 32 {
		panic("AES KEY IS NOT MATCH PLEASE USING 16,24,32")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.GetConfigConnection().Database.CRM.Username, c.GetConfigConnection().Database.CRM.Password, c.GetConfigConnection().Database.CRM.Host, c.GetConfigConnection().Database.CRM.Port, c.GetConfigConnection().Database.CRM.Database)
	if isuat {
		return
	}
	if c.GetConfigConnection().Database.CRM.Enable == true {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		if err != nil {
			panic("Failed to connect to the database: " + err.Error())
		}
		c.database = db
	} else {
		panic("Database Not Enable: ")
	}
}

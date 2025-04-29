package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	dsn := Config.GetString("database.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

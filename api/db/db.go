package db

import (
	"tara/api/config"
	"tara/api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.Cfg.DBURI)
	if err != nil {
		return nil, err
	}

	DB = db

	db.AutoMigrate(&models.Users{})

	return db, err
}

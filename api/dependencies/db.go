package dependencies

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tara/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DBURI)
	if err != nil {
		return nil, err
	}

	DB = db

	db.AutoMigrate(&models.Users{})

	return db, err
}
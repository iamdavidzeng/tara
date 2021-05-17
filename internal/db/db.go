package db

import (
	"github.com/iamdavidzeng/tara/api/models"
	"github.com/iamdavidzeng/tara/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Storage *gorm.DB
}

var D *Database = &Database{}

func (db *Database) Init() error {
	conn, err := gorm.Open(mysql.Open(config.Cfg.DB.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	conn.AutoMigrate(&models.Users{}, &models.Posts{})

	db.Storage = conn

	return nil
}

package api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() (*gorm.DB, error) {
	client, err := gorm.Open(mysql.Open(Cfg.MysqlCfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client.AutoMigrate(&Users{}, &Posts{})

	DB = client

	return client, nil
}

package config

import (
	"phone_review/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	//declare struct config & variable connectionString
	dsn := "admin:root@tcp(127.0.0.1:3306)/phone_review?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

var (
	DB *gorm.DB
)

func initMigration() {
	DB.AutoMigrate(&db.Admins{},
		&db.Advances{},
		&db.Audios{},
		&db.Batteries{},
		&db.Desains{},
		&db.Descs{},
		&db.Displays{},
		&db.Merks{},
		&db.Types{},
		&db.Performs{},
		&db.Cameras{},
		&db.Features{})
}

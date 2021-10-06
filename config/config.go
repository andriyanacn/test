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
	DB.AutoMigrate(&db.Admins{})
	DB.AutoMigrate(&db.Advances{})
	DB.AutoMigrate(&db.Audios{})
	DB.AutoMigrate(&db.Batteries{})
	DB.AutoMigrate(&db.Desains{})
	DB.AutoMigrate(&db.Descs{})
	DB.AutoMigrate(&db.Displays{})
	DB.AutoMigrate(&db.Merks{})
	DB.AutoMigrate(&db.Types{})
	DB.AutoMigrate(&db.Performs{})
	DB.AutoMigrate(&db.Cameras{})
	DB.AutoMigrate(&db.Features{})
}

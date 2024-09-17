package config

import (
	"Learn-Gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/learngin?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed conection to database")
	}

	DB = database

	//migrate schema
	database.AutoMigrate(&models.Article{})

	database.AutoMigrate(&models.User{}, &models.Article{})

}

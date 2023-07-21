package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/phone_go"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Phone{})

	DB = database
}

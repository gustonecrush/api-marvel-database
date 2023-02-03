// this file is used to settup the connection to the database
package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/api_marvel"))
	if err != nil {
		panic(err)
	}

	// running automigrate
	database.AutoMigrate(&Hero{})
	database.AutoMigrate(&User{})

	DB = database
}
// CREATE CONNECTION TO DATABASE

package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:bagaskaramadhan97@tcp(localhost:3306)/go_toko"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("OK")
	}

	database.AutoMigrate(&Product{})
	
	DB = database
}

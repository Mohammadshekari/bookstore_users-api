package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Connect() {
	connection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("ERROR => ", err)
	}
	fmt.Println("DB => ", Db)
	Db = connection
}

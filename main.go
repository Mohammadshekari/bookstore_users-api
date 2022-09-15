package main

import (
	"fmt"
	"github.com/mohammadshekari/bookstore_users-api/app"
	"github.com/mohammadshekari/bookstore_users-api/domain/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("SQL RUN FAILED")
		return
	}
	db.AutoMigrate(&users.User{})
	app.StartApplication()
}

package main

import (
	"github.com/mohammadshekari/bookstore_users-api/app"
	"github.com/mohammadshekari/bookstore_users-api/database"
	"github.com/mohammadshekari/bookstore_users-api/domain/users"
)

func main() {
	database.Connect()
	database.Db.AutoMigrate(&users.User{})
	app.StartApplication()
}

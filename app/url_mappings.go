package app

import (
	"github.com/mohammadshekari/bookstore_users-api/controllers/ping"
	"github.com/mohammadshekari/bookstore_users-api/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.POST("/users", users.CreateUser)
}

package app

import (
	"github.com/yftzz/golang-bookstore/bookstore_users-api/controllers/pingController"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/controllers/usersController"
)

func mapUrls() {
	router.GET("/ping", pingController.Ping)

	router.GET("/users/:user_id", usersController.GetUser)
	router.POST("/users", usersController.CreateUser)
}

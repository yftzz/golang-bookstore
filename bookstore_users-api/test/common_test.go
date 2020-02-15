package main

import (
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/controllers/pingController"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/controllers/usersController"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pingController.Ping)
	router.POST("/users", usersController.CreateUser)
	router.GET("/users/:user_id", usersController.GetUser)
	return router
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("incorrect status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if strings.Index(got, want) < 0 {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

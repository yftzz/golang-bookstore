package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func TestController(t *testing.T) {
	router = getRouter()

	t.Run("test ping controller and its return value", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assertStatus(t, w.Code, http.StatusOK)
		assertResponseBody(t, w.Body.String(), "pong")
	})

	t.Run("test invalid user creation", func(t *testing.T) {
		userString := `{"id":0,"first_name":"Yifei","last_name":"Tang","email":"yifei.tang@icloud.com","date_created":""}`

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(userString)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusBadRequest)
		assertResponseBody(t, w.Body.String(), "Invalid user ID")
	})

}

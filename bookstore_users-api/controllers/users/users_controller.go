package users

import (
	"net/http"

	"github.com/yftzz/golang-bookstore/bookstore_users-api/services"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/utils/errors"

	"github.com/yftzz/golang-bookstore/bookstore_users-api/domain/users"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "under construction")
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, res)
}

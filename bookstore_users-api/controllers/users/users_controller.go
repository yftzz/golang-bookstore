package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yftzz/golang-bookstore/bookstore_users-api/services"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/utils/errors"

	"github.com/yftzz/golang-bookstore/bookstore_users-api/domain/users"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	fmt.Println("user id is ", userID)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user ID")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println("creating user")
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, res)
}

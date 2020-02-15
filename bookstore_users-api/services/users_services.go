package services

import (
	"github.com/yftzz/golang-bookstore/bookstore_users-api/domain/users"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	res := users.User{ID: userID}
	if err := res.Get(); err != nil {
		return nil, err
	}
	return &res, nil
}
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

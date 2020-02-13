package users

import (
	"fmt"

	"github.com/yftzz/golang-bookstore/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	res := usersDB[user.ID]

	if res == nil {
		return errors.NewBadFoundError(fmt.Sprintf("User %d not found", user.ID))
	}

	user.ID = res.ID
	user.Email = res.Email
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.DateCreated = res.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	currUser := usersDB[user.ID]
	if currUser != nil {
		if currUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already registered",
				user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists in DB",
			user.ID))
	}

	usersDB[user.ID] = user
	return nil
}

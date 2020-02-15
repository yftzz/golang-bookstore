package users

import (
	"fmt"
	"strings"

	"github.com/yftzz/golang-bookstore/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/utils/errors"
	"github.com/yftzz/golang-bookstore/bookstore_users-api/utils/time_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, " +
		"email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = time_utils.GetNowString()
	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists",
				user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error trying to save user: %s",
			err.Error()))
	}
	userID, err := insertRes.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error trying to save user: %s",
			err.Error()))
	}
	user.ID = userID
	return nil
}

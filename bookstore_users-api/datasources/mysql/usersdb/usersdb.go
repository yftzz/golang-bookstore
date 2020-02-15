package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = "tyf19981211"
	mysqlUsersHost     = "localhost:3306"
	mysqlUsersSchema   = "users_db"
)

var (
	Client *sql.DB
)

// create connection to the driver
func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysqlUsersUsername,
		mysqlUsersPassword,
		mysqlUsersHost,
		mysqlUsersSchema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")
}

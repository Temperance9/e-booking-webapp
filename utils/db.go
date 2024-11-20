package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/ebooking")
	if err != nil {
		panic(err.Error())
	}

	return db, err
}

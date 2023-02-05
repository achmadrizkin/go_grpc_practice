package database

import (
	"database/sql"
	"fmt"

	"github.com/achmadrizkin/go_grpc_practice/config"
)

// DB is database
var DB *sql.DB

func Connect() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected to the database")
	return nil
}

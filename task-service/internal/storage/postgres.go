package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	dsn := "user=admin dbname=gotest sslmode=disable password=1234"
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

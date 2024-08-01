package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
  var err error
  DB, err = sql.Open("postgres", "user=postgres dbname=scti-db sslmode=disable")
  if err != nil {
    return err
  }
  return nil
}

func CloseDatabase() error {
  return DB.Close()
}

func CreateUser() error {
  return nil
}

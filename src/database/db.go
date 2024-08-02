package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
    host     = "127.0.0.1"
    port     = 5432
    user     = "postgres"
    password = "root"
    dbname   = "scti-db"
)


func OpenDatabase() error {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

  var err error
  DB, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    return err
  }
  
  err = DB.Ping()
  if err != nil {
       log.Fatal(err)
  }
  return nil
}

func CloseDatabase() error {
  return DB.Close()
}

func CreateUser(Email string, hash string) error {
    query := `
        INSERT INTO users (email)
        VALUES ($1)
        RETURNING id
    `
    var userID int
    err := DB.QueryRow(query, Email).Scan(&userID)
    if err != nil {
        return fmt.Errorf("não foi possível inserir o usuário: %v", err)
    }

    queryPasswd := `
	INSERT INTO passwd (id, passwd)
	VALUES ($1, $2)
    `
    tx, err := DB.Begin()
    if err != nil {
        log.Fatal(err)
    }

    _, err = tx.Exec(queryPasswd, userID, hash)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("não foi possível inserir a senha: %v", err)
    }

    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("não foi possível confirmar o signup: %v", err)
    }

    fmt.Printf("Novo usuário inserido com ID: %d\n", userID)
    return nil
}

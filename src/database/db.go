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
  tx, err := DB.Begin()
  if err != nil {
    log.Fatal(err)
  }

  query := `
  INSERT INTO users (email)
  VALUES ($1)
  RETURNING id
  `

  var userID int
  err = tx.QueryRow(query, Email).Scan(&userID)
  if err != nil {
    tx.Rollback()
    return fmt.Errorf("não foi possível inserir o usuário: %v", err)
  }

  queryPasswd := `
  INSERT INTO passwd (id, passwd)
  VALUES ($1, $2)
  `

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

func UserExists(email string) (bool, error) {
    query := `
    SELECT EXISTS(
        SELECT 1 FROM users WHERE email = $1
    )
    `

    var exists bool
    err := DB.QueryRow(query, email).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("could not check if user exists: %v", err)
    }

    return exists, nil
}

func GetHash(email string) (string) {
    query := `
    SELECT passwd.passwd
    FROM users
    JOIN passwd ON users.id = passwd.id
    WHERE users.email = $1
    `

    var hash string
    err := DB.QueryRow(query, email).Scan(&hash)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Printf("no user found with email: %s\n", email)
            return ""
        }
        fmt.Printf("could not retrieve password hash: %v\n", err)
        return ""
    }

    return hash
}

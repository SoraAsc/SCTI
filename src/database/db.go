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

func CreateUser(Email string, hash string, UUIDString string) error {
  tx, err := DB.Begin()
  if err != nil {
    log.Fatal(err)
  }

  query := `
  INSERT INTO users (email, uuid, verificationCode)
  VALUES ($1, $2, $3)
  RETURNING id
  `

  var userID int
  err = tx.QueryRow(query, Email, UUIDString, UUIDString[:5]).Scan(&userID)
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

func GetId(uuid string) (int, error) {
    query := `
    SELECT id
    FROM users
    WHERE users.uuid = $1
    `

    var id int
    err := DB.QueryRow(query, uuid).Scan(&id)
    if err != nil {
        if err == sql.ErrNoRows {
            return -1, fmt.Errorf("no user found with uuid: %s\n", uuid)
        }
        return -1, fmt.Errorf("could not retrieve id: %v\n", err)
    }

    return id, nil
}

func GetCode(uuid string) (string, error) {
    query := `
    SELECT verificationCode
    FROM users
    WHERE users.uuid = $1
    `

    var code string
    err := DB.QueryRow(query, uuid).Scan(&code)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", fmt.Errorf("no user found with uuid: %s\n", uuid)
        }
        return "", fmt.Errorf("could not retrieve id: %v\n", err)
    }

    return code, nil
}

func GetUUID(email string) (string) {
    query := `
    SELECT uuid
    FROM users
    WHERE users.email = $1
    `

    var uuid string
    err := DB.QueryRow(query, email).Scan(&uuid)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Printf("no user found with email: %s\n", email)
            return ""
        }
        fmt.Printf("could not retrieve uuid: %v\n", err)
        return ""
    }

    return uuid
}

func GetStanding(email string) (bool) {
    query := `
    SELECT isVerified
    FROM users
    WHERE users.email = $1
    `

    var accStatus bool
    err := DB.QueryRow(query, email).Scan(&accStatus)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Printf("no user found with email: %s\n", email)
            return false
        }
        fmt.Printf("could not retrieve id: %v\n", err)
        return false
    }

    return accStatus
}

func SetStanding(uuid string, standing bool) (error) {
  tx, err := DB.Begin()
  if err != nil {
    log.Fatal(err)
  }

  query := `
  UPDATE users 
  SET isVerified = $1
  WHERE uuid = $2
  `

  _, err = tx.Exec(query, standing, uuid)
  if err != nil {
    tx.Rollback()
    return fmt.Errorf("não foi possível verificar o usuário: %v", err)
  }

  err = tx.Commit()
  if err != nil {
    return fmt.Errorf("não foi possível confirmar a transação de verificação: %v", err)
  }

  fmt.Println("Usuário verificado")
  return nil
}

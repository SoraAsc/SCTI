package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func OpenDatabase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWD"),
		os.Getenv("DB_NAME"),
	)

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

func CreateUser(Email string, hash string, UUIDString string, name string) error {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	query := `
  INSERT INTO users (email, uuid, verificationCode, name)
  VALUES ($1, $2, $3, $4)
  RETURNING id
  `

	var userID int
	err = tx.QueryRow(query, Email, UUIDString, UUIDString[:5], name).Scan(&userID)
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

func GetHash(email string) string {
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

func GetCodeByEmail(email string) (string, error) {
	query := `
  SELECT verificationCode
  FROM users
  WHERE users.email = $1
  `

	var code string
	err := DB.QueryRow(query, email).Scan(&code)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with email: %s\n", email)
		}
		return "", fmt.Errorf("could not retrieve id: %v\n", err)
	}

	return code, nil
}

func GetUUID(email string) string {
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

func GetEmail(uuid string) string {
	query := `
  SELECT email
  FROM users
  WHERE users.uuid = $1
  `

	var email string
	err := DB.QueryRow(query, uuid).Scan(&uuid)
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

func GetStanding(email string) bool {
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

func SetStanding(uuid string, standing bool) error {
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

func GetAdmin(uuid string) bool {
	query := `
  SELECT isAdmin
  FROM users
  WHERE users.uuid = $1
  `

	var admStatus bool
	err := DB.QueryRow(query, uuid).Scan(&admStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("no user found with uuid: %s\n", uuid)
			return false
		}
		fmt.Printf("could not retrieve admin status: %v\n", err)
		return false
	}

	return admStatus
}

func SetAdmin(uuid string, admStatus bool) error {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	query := `
  UPDATE users 
  SET isAdmin = $1
  WHERE uuid = $2
  `

	_, err = tx.Exec(query, admStatus, uuid)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Usuário inexistente: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("não foi possível confirmar a transação: %v", err)
	}

	return nil
}

func DeleteUser(userUUID string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
  SELECT id FROM users
  WHERE uuid = $1
  `

	var userID int
	err = tx.QueryRow(query, userUUID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	query = `
  DELETE FROM credits
  WHERE purchase_id IN(
  SELECT purchase_id
  FROM purchases
  WHERE user_id = $1
  )
  `

	_, err = tx.Exec(query, userUUID)
	if err != nil {
		return err
	}

	query = `
  DELETE FROM purchases
  WHERE user_id = $1
  `

	_, err = tx.Exec(query, userUUID)
	if err != nil {
		return err
	}

	query = `
  DELETE FROM registrations
  WHERE user_id = $1
  `

	_, err = tx.Exec(query, userUUID)
	if err != nil {
		return err
	}

	query = `
  DELETE FROM passwd
  WHERE id = $1
  `

	_, err = tx.Exec(query, userID)
	if err != nil {
		return err
	}

	query = `
  DELETE FROM users
  WHERE uuid = $1
  `

	_, err = tx.Exec(query, userUUID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func ChangeUserPassword(userUUID string, newPassword string) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	var userID int
	err = tx.QueryRow("SELECT id FROM users WHERE uuid = $1", userUUID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no user found with UUID: %s", userUUID)
		}
		return fmt.Errorf("failed to fetch user: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	_, err = tx.Exec("UPDATE passwd SET passwd = $1 WHERE id = $2", string(hashedPassword), userID)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func MarkAsPaid(email string) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	query := `
  UPDATE users
  SET isPaid = TRUE
  WHERE email = $1
  `

	result, err := tx.Exec(query, email)
	if err != nil {
		return fmt.Errorf("failed to update user payment status: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with email: %s", email)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func IsUserPaid(uuid string) (bool, error) {
	var isPaid bool

	query := `
  SELECT isPaid
  FROM users
  WHERE uuid = $1
  `
	err := DB.QueryRow(query, uuid).Scan(&isPaid)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("no user found with UUID: %s", uuid)
		}
		return false, fmt.Errorf("failed to fetch user payment status: %v", err)
	}

	return isPaid, nil
}

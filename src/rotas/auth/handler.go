package auth

import (
  "bufio"
  "fmt"
  "log"
  "net/http"
  "os"
  "strings"

  "golang.org/x/crypto/bcrypt"
)

type Handler struct{}

type User struct {
  Email string
  Password string 
}

func UserExists(Email string)(userExists bool) {
  file, err := os.Open("passwords.txt")
  if err != nil && !os.IsNotExist(err) {
    log.Fatal(err)
  }
  defer file.Close()

  if file != nil {
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      line := scanner.Text()
      parts := strings.SplitN(line, ":", 3)
      if len(parts) == 3 && parts[0] == Email {
        userExists = true
        break
      }
    }
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
  }
  return userExists
}


func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}



func (h *Handler) GetPrintCookies(w http.ResponseWriter, req *http.Request) {
  var returnStr string
  for _, cookie := range req.Cookies() {
    returnStr = returnStr + cookie.Name + ":" + cookie.Value + "\n"
  }
  fmt.Fprint(w, returnStr)
}

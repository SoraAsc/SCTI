package auth

import (
  "SCTI/fileserver"
  "bufio"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
  "strings"

  "golang.org/x/crypto/bcrypt"
  "github.com/google/uuid"
)

type Handler struct{}

func (h *Handler) GetSignup(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/login.gohtml")
  t.Execute(w, nil)
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

func VerifyLogin(user User)(login bool) {
  file, err := os.Open("passwords.txt")
  if err != nil && !os.IsNotExist(err) {
    log.Fatal(err)
  }
  defer file.Close()

  var storedHash string
  var found bool

  if file == nil {
    return false
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.SplitN(line, ":", 3)
    if len(parts) == 3 && parts[0] == user.Email {
      storedHash = parts[1]
      found = true
      break
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  if !found {
    println("Verify Password: User not found")
    return false
  } else {
    println("Verify Pasword: User found")
  }

  if CheckPasswordHash(user.Password, storedHash) {
    login = true
  } else {
    login = false
  }
  return login
}

type User struct {
  Email string
  Password string 
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
  println("In PostSignup")

  var user User

  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }

  if UserExists(user.Email) {
    println("User already exists")
    return
  }

  hash, _ := HashPassword(user.Password)

  file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  _, err = file.WriteString(fmt.Sprintf("%s:%s:%s\n", user.Email, hash, uuid.NewString()))
  if err != nil {
    log.Fatal(err)
    return
  }

  fmt.Println("E-Mail: ", user.Email)
  fmt.Println("Password: ", user.Password)
  fmt.Println("Hash: ", hash)
}

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
  println("In PostLogin")

  var user User

  if r.Header.Get("Content-type") == "application/json" {
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if err := r.ParseForm(); err != nil {
      fmt.Println("r.Form dentro if: ", r.Form)
      log.Fatal(err)
    }
    user.Email = r.FormValue("Email")
    user.Password = r.FormValue("Senha")
  }
  if VerifyLogin(user) {
    println("Successful Login")
  } else {
    println("Login Failed")
  }
}

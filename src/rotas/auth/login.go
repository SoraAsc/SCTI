package auth

import (
  "SCTI/fileserver"
  "encoding/json"
  "net/http"
  "strings"
  "bufio"
  "time"
  "log"
  "fmt"
  "os"
)

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("accessToken")
  if err == nil && cookie != nil {
    http.Redirect(w, r, "/cookies", http.StatusSeeOther)
    return
  }

  var t = fileserver.Execute("template/login.gohtml")
  t.Execute(w, nil)
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
  if VerifyLogin(user, w) {
    println("Successful Login")
  } else {
    println("Login Failed")
  }
}

func VerifyLogin(user User, w http.ResponseWriter)(login bool) {
  file, err := os.Open("passwords.txt")
  if err != nil && !os.IsNotExist(err) {
    log.Fatal(err)
  }
  if file == nil {
    return false
  }
  defer file.Close()

  var storedHash string
  var found bool
  var uuid string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.SplitN(line, ":", 3)
    if len(parts) == 3 && parts[0] == user.Email {
      storedHash = parts[1]
      uuid = parts[2]
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
    cookie := http.Cookie{
      Name: "accessToken",
      Value: uuid,
      Expires: time.Now().Add(2 * 24 * time.Hour),
      Secure: false,
      HttpOnly: true,
      Path: "/",
    }
    http.SetCookie(w, &cookie)
  } else {
    login = false
  }
  return login
}



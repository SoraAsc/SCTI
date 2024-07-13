package auth

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "SCTI/fileserver"
  "github.com/lengzuo/supa"
  "github.com/lengzuo/supa/dto"
)

type Handler struct{
  S *supabase.Client
}

type User struct {
  Email string
  Password string 
}

func (h *Handler) PostSignup(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
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
    user.Email = r.FormValue("Nome")
    user.Password = r.FormValue("Idade")
  }

  fmt.Println(user.Email)
  fmt.Println(user.Password)

  body := dto.SignUpRequest{
    Email:    user.Email,
    Password: user.Password,
  }

  _, err := h.S.Auth.SignUp(ctx, body)
  if err == nil {
    panic("Panicked at PostSignup")
  } else {
    fmt.Printf("We signed up?")
  }
}

func (h *Handler) GetSignup(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/signup.gohtml")
  t.Execute(w, nil)
}

func (h *Handler) GetLogin(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Pagina de login")
}

func RegisterRoutes(mux *http.ServeMux, s *supabase.Client) {
  handler := &Handler{S: s}
  mux.HandleFunc("GET /signup", handler.GetSignup)
  mux.HandleFunc("POST /signup", handler.GetSignup)
  mux.HandleFunc("/login", handler.GetLogin)
}

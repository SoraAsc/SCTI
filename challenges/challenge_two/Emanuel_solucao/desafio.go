package main

import (
  "encoding/json"
  "fmt"
  "strconv"
  "log"
  "net/http"
)


type Pessoa struct {
    Nome string
    Idade int
    Cadastrado bool
    Id int
}


func Create(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
     http.ServeFile(w, r, "form.html")
  case "POST":
    var PostReceive Pessoa

    if r.Header.Get("Content-type") == "application/json" {
      err := json.NewDecoder(r.Body).Decode(&PostReceive)
      if err != nil {
        log.Fatal(err)
      }
    } else {
      // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
      if err := r.ParseForm(); err != nil {
        fmt.Println("r.Form dentro if: ", r.Form)
        log.Fatal(err)
      }
      PostReceive.Nome = r.FormValue("Nome")
      PostReceive.Idade, _ = strconv.Atoi(r.FormValue("Idade"))
      PostReceive.Cadastrado, _ = strconv.ParseBool(r.FormValue("Cadastrado"))
    }

    PostReceive.Id = NextId
    NextId++
    Pessoas = append(Pessoas, PostReceive)
  default:
    fmt.Fprintf(w, "Apenas GET e POST é suportado.")
  }


}


func List(w http.ResponseWriter, _ *http.Request) {
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")
  jsonResp, err := json.Marshal(Pessoas)
  if err != nil {
    fmt.Println("Ocorreu um erro com o JSON marshal. Err: ", err)
  }
  w.Write(jsonResp)
}


func Registered(w http.ResponseWriter, _ *http.Request) {
  var Cadastrados []Pessoa
  for _, Pessoa := range Pessoas{
    if Pessoa.Cadastrado {
      Cadastrados = append(Cadastrados, Pessoa)
    }
  }


  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")
  jsonResp, err := json.Marshal(Cadastrados)
  if err != nil {
    fmt.Println("Ocorreu um erro com o JSON marshal. Erro: ", err)
  }
  w.Write(jsonResp)
}


var Pessoas = []Pessoa{{Nome: "Sophia", Idade: 21, Cadastrado: true, Id: 10}, {Nome: "Teste", Idade: 32, Cadastrado: false, Id: 11}}
var NextId = 12


func main(){
  http.Handle("/", http.FileServer(http.Dir("./"))) //home page
  http.HandleFunc("/app/user/create-user", Create)
  http.HandleFunc("/app/user/list-users", List)
  http.HandleFunc("/app/user/registered-users", Registered)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

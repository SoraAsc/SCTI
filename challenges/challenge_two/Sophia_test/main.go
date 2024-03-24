package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id         int
	Nome       string
	Idade      int
	Cadastrado bool
}

func (u User) String() string {
	return fmt.Sprintf("%v: %v tem %v anos\nstatus do cadastro (%v)", u.Id, u.Nome, u.Idade, u.Cadastrado)
}

var Users = []User{
	{Id: 10, Nome: "Sophia", Idade: 21, Cadastrado: true},
	{Id: 11, Nome: "Teste", Idade: 32, Cadastrado: false},
}

func main() {
	http.HandleFunc("/app/user/create-user", CreateUser)
	http.HandleFunc("/app/user/list-users", ListUsers)
	http.HandleFunc("/app/user/registered-users", RegisteredUsers)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user.Id = 12
		fmt.Println(user)
		Users = append(Users, user)
		w.WriteHeader(http.StatusCreated)
	}
	fmt.Println()
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	method := r.Method
	if method == "GET" {
		resp, err := json.Marshal(Users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write(resp)
	}
	fmt.Println()
}

func RegisteredUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	method := r.Method

	var registered []User

	if method == "GET" {
		for _, u := range Users {
			if u.Cadastrado {
				registered = append(registered, u)
			}
		}

		resp, err := json.Marshal(registered)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write(resp)
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var users = map[string]User{}

func InsertIntoDB(r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")
	users[username] = User{Username: username, Password: password}
	return true
}

func GetFromDB() {
	fmt.Println("checkDB")
}

func Authenticate() (loggedIn bool) {
	fmt.Println("authenticate")
	loggedIn = true
	return
}

package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Username = "Nathan"
var LoggedIn = false

var currentPage string = "/"

type UserSession struct {
	Username    string
	LoggedIn    bool
	Reading     Reading
	Category    Category
	Categories  []Category
	ReadingList []ReadingList
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	currentPage = r.URL.Path
	var data UserSession
	data.Username = Username
	data.LoggedIn = LoggedIn

	if InsertIntoDB(r) {
		generateHTML(w, data, "signup", "navbar", "footer", "signup")
	} else {
		//insert failed
		generateHTML(w, data, "signup", "navbar", "footer", "signup")
	}
}

func login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentPage = r.URL.Path
	print("login\n")
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.Username = Username
	data.LoggedIn = LoggedIn

	generateHTML(w, data, "login", "navbar", "footer", "login")
}

func logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	fmt.Printf("URL path : %s\n", r.URL.Path)
	LoggedIn = false

	if currentPage == "/account" {
		currentPage = "/"
	}
	http.Redirect(w, r, currentPage, 302)
}

func account(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	currentPage = r.URL.Path
	//username := r.FormValue("username")
	//_, ok := users[username]
	//if !ok {
	if !LoggedIn {
		print("not logged in\n")
		http.Redirect(w, r, "/login", 302)
		return
	}
	print("going to account\n")
	var data UserSession
	data.Username = Username
	data.LoggedIn = LoggedIn
	generateHTML(w, data, "account", "navbar", "footer", "account")
}

func signup_account(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	print("signup_account" + p.ByName("name"))
	currentPage = r.URL.Path

	LoggedIn = Authenticate()
	//if true
	http.Redirect(w, r, "/", 302)

}

func authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("authenticate: %s\n", r.FormValue("username"))
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	currentPage = r.URL.Path

	LoggedIn = Authenticate()
	//if true
	http.Redirect(w, r, "/", 302)
}

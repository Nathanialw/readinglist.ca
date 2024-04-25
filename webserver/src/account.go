package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var currentPage string = "/"

type UserSession struct {
	Username         string
	Admin            bool
	LoggedIn         bool
	Reading          Reading
	Book             Book
	Category         Category
	Categories       []Category
	ReadingList      []ReadingList
	Books            []Book
	JsonBooks        string
	JsonReadingLists string

	FavouriteBooks     []Book
	JsonFavouriteBooks string
	QueuedBooks        []Book
	JsonQueuedBooks    string
	ReadingHistory     []Book
	JsonReadingHistory string
	//post history
	//post history json
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	currentPage = r.URL.Path
	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	generateHTML(w, data, "signup", "navbar", "footer", "signup")
}

func login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentPage = r.URL.Path
	print("login\n")
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	generateHTML(w, data, "login", "navbar", "footer", "login")
}

func logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	fmt.Printf("URL path : %s\n", r.URL.Path)

	ClearAllCookies(w)

	if currentPage == "/account" {
		currentPage = "/"
	}
	http.Redirect(w, r, currentPage, 302)
}

func account(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	currentPage = r.URL.Path

	if !LoginStatus(r) {
		print("not logged in\n")
		http.Redirect(w, r, "/login", 302)
		return
	}

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Username = GetUsername(r)
	//can save a hashed value as a cookie and use that to get the user's data

	data.FavouriteBooks, _ = GetFavouriteBooks(data.Username)
	data.QueuedBooks, _ = GetQueuedBooks(data.Username)
	data.ReadingHistory, _ = GetReadingHistory(data.Username)
	data.ReadingList, _ = GetSavedReadingLists(data.Username)
	//get post history

	//stringify and send and a json for the js to handle
	generateHTML(w, data, "account", "navbar", "footer", "account")
}

func signup_account(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	print("signup_account\n" + p.ByName("name"))

	if InsertIntoDB(r) {
		SetCookies(w, r)
		http.Redirect(w, r, "/", 302)
	}
	http.Redirect(w, r, "/signup", 302)
}

func authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("authenticate: %s\n", r.FormValue("username"))
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	if Authenticate(r) {
		SetCookies(w, r)
		http.Redirect(w, r, "/", 302)
	}
	http.Redirect(w, r, "/login", 302)
}

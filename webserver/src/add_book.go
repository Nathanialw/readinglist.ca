package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func addbook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	generateHTML(w, data, "addbook", "navbar", "footer", "addbook")
}

func submitbook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	//check on the browser side first and give feedback to the user
	//then check on the server side
	//if this fails it should redirect to the addbook page

	var data UserSession
	data.LoggedIn = LoginStatus(r)

	if !data.LoggedIn {
		fmt.Println("not logged in")
		notfound(w, r, p)
		return
	}

	VerifyAndInsertBook(w, r, contentDB)
	http.Redirect(w, r, "/addbook", http.StatusSeeOther)
}

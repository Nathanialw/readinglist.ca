package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ShowPostHistory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)

	generateHTML(w, data, "posthistory", "posthistory")
}

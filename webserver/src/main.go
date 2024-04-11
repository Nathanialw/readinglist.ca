// to build
// go build -o ../app/main  main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	Init()
	r := httprouter.New()
	r.NotFound = http.StripPrefix("/", http.FileServer(http.Dir("../../public/")))

	r.GET("/404", notfound)

	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/about", about)
	r.GET("/category/*categoryPath", category)
	r.GET("/readinglist/*listPath", readinglist)

	r.GET("/signup", signup)
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/account", account)
	r.POST("/signup_account", signup_account)
	r.POST("/authenticate", authenticate)

	server := http.Server{
		Addr:    "localhost:12001",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func notfound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)

	generateHTML(w, data, "notfound", "navbar", "footer", "notfound")
}

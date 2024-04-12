// to build
// go build -o ../app/main  main.go

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	Init()
	r := httprouter.New()

	//r.NotFound = http.StripPrefix("/", http.FileServer(http.Dir("../../public/")))
	fs := http.FileServer(http.Dir("../../public/"))
	r.NotFound = http.StripPrefix("/", fileServerWith404(fs))

	r.GET("/notfound", notfound)

	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/about", about)
	r.GET("/category/*categoryPath", category)
	r.GET("/readinglist/*listPath", readinglist)

	r.GET("/admin", admin)
	r.GET("/signup", signup)
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/account", account)
	r.POST("/signup_account", signup_account)
	r.POST("/authenticate", authenticate)

	r.GET("/addbook", addbook)
	r.GET("/addreadinglist", addreadinglist)
	r.POST("/submitbook", submitbook)

	r.GET("/readinglists", underconstruction)
	r.GET("/favoritedbooks", underconstruction)
	r.GET("/readinghistory", underconstruction)
	r.GET("/queuedbooks", underconstruction)
	r.GET("/posthistory", underconstruction)

	server := http.Server{
		Addr:    "localhost:12001",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func fileServerWith404(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := os.Stat("../../public/" + r.URL.Path)

		if os.IsNotExist(err) {
			fmt.Printf("file does not exist: %s\n", r.URL.Path)
			// If the file does not exist, serve your 404 page
			notfound(w, r, httprouter.Params{})
			return
		}

		fmt.Printf("file exists: %s\n", r.URL.Path)
		// If the file exists, serve it
		h.ServeHTTP(w, r)
	}
}

func notfound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	generateHTML(w, data, "notfound", "navbar", "footer", "notfound")
}

func underconstruction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	generateHTML(w, data, "undercontruction", "navbar", "footer", "undercontruction")
}

func admin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	if !data.Admin {
		notfound(w, r, p)
		return
	}
	generateHTML(w, data, "admin", "navbar", "footer", "admin")
}

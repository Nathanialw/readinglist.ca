// to build
// go build -o ../app/main  main.go

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PageData struct {
	Title string
	Body  string
}

func main() {
	r := httprouter.New()
	// r.GET("/goapp", homeHandler)
	// r.OPTIONS("/goapp", handleOptions) // Add this line

	r.NotFound = http.StripPrefix("/", http.FileServer(http.Dir("../../public/")))

	server := http.Server{
		Addr:    "localhost:12001",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

// func handleOptions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 	w.WriteHeader(http.StatusNoContent)
// }

func homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
	}
	tmpl, err := template.ParseFiles("../templates/charts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	// Handle requests for individual products here
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Handle requests for the user API here
}

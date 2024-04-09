// to build
// go build -o ../app/main  main.go

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type PageData struct {
	Title string
	Body  string
}

func main() {
	Init()
	r := httprouter.New()
	r.NotFound = http.StripPrefix("/", http.FileServer(http.Dir("../../public/")))

	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/about", about)
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/account", account)

	r.GET("/category/*categoryPath", category)
	r.GET("/readinglist/*listPath", readinglist)

	server := http.Server{
		Addr:    "localhost:12001",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, name string, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("../templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, name, data)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}

func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	data, _ := Categories()

	generateHTML(w, data, "landing", "navbar", "footer", "dailylist", "category", "landing")
}

func contact(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
	}

	generateHTML(w, data, "contact", "navbar", "footer", "contact")
}

func about(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
	}

	generateHTML(w, data, "about", "navbar", "footer", "about")
}

func login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
	}

	generateHTML(w, data, "login", "navbar", "footer", "login")
}

func logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
	}

	generateHTML(w, data, "logout", "navbar", "footer", "logout")
}

func account(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
	}

	generateHTML(w, data, "account", "navbar", "footer", "account")
}

func category(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	//strip off the end of the url
	list := strings.TrimPrefix(p.ByName("categoryPath"), "/")
	fmt.Printf("category: %s\n", list)
	data, _ := ReadingLists(list)

	generateHTML(w, data, "category", "navbar", "footer", "category")
}

func readinglist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	list := strings.TrimPrefix(p.ByName("listPath"), "/")
	fmt.Printf("readinglist: %s\n", list)
	var data Reading
	data.Books, _ = Books(list)
	data.Reading_list, _ = GetReadingList(list)

	generateHTML(w, data, "readinglist", "navbar", "footer", "readinglist")
}

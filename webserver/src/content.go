package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

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
	currentPage = r.URL.Path
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	var data UserSession
	data.Categories, _ = Categories()
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	generateHTML(w, data, "landing", "navbar", "footer", "dailylist", "category", "landing")
}

func contact(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentPage = r.URL.Path
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	generateHTML(w, data, "contact", "navbar", "footer", "contact")
}

func about(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentPage = r.URL.Path
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	generateHTML(w, data, "about", "navbar", "footer", "about")
}

func category(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentPage = r.URL.Path
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	//strip off the end of the url
	list := strings.TrimPrefix(p.ByName("categoryPath"), "/")
	fmt.Printf("category: %s\n", list)
	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Category, _ = GetCategory(list)
	data.ReadingList, _ = ReadingLists(list)

	generateHTML(w, data, "category", "navbar", "footer", "category")
}

func readinglist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentPage = r.URL.Path
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	list := strings.TrimPrefix(p.ByName("listPath"), "/")
	fmt.Printf("readinglist: %s\n", list)
	//var data Reading
	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Reading.Books, _ = Books(list)
	data.Reading.Reading_list, _ = GetReadingList(list)

	generateHTML(w, data, "readinglist", "navbar", "footer", "readinglist")
}

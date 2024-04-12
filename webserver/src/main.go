// to build
// go build -o ../app/main  main.go

package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

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

func addbook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	generateHTML(w, data, "addbook", "navbar", "footer", "addbook")
}

func addreadinglist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	generateHTML(w, data, "addreadinglist", "navbar", "footer", "addreadinglist")
}

func textNotEmpty(title string) bool {
	if title == "" {
		fmt.Println("title is empty")
		return false
	}
	return true
}

func validYear(year string, era string) bool {
	if year == "" || era == "" {
		fmt.Println("year is empty")
		return false
	}
	for _, c := range year {
		if c < '0' || c > '9' {
			fmt.Println("year is not a number")
			return false
		}
	}
	if era != "BC" && era != "AD" {
		fmt.Println("era is not BC or AD")
		return false
	}
	if len(year) > 4 {
		fmt.Println("year is more than 4 characters")
		return false
	}
	if year == "0" {
		fmt.Println("year is no year 0 BC or 0 AD")
		return false
	}
	if era == "AD" && year > "2024" {
		fmt.Println("the max year is 2024 AD")
		return false
	}
	return true
}

func verifyImage(handler *multipart.FileHeader) bool {
	if handler.Filename == "" {
		fmt.Println("image is empty")
		return false
	}
	// Check the file type
	fileType := handler.Header.Get("Content-Type")
	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		break
	default:
		fmt.Println("File is not an image")
		return false
	}
	if handler.Size > 4*1024*1024 {
		fmt.Printf("File is too large (max 4MB) %d\n", handler.Size)
		return false
	}
	//check if the file is an image
	if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
		fmt.Printf("file is not an image: %s\n", handler.Header.Get("Content-Type"))
		return false
	}
	//check if the file already exists
	if _, err := os.Stat("/assets/images/book_covers/" + handler.Filename); err == nil {
		//maybe append a number to the filename?
		fmt.Printf("file already exists: %s\n", handler.Filename)
		return false
	}

	return true
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

	//verify book data is formatted correctly
	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	author := r.FormValue("author")
	publish_year := r.FormValue("publish_year")
	publish_era := r.FormValue("publish_era")
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	// Retrieve the file from form data
	file, handler, err := r.FormFile("image") // "image" is the name of the file input field
	if err != nil {
		fmt.Printf("error retrieving the file: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	synopsis := r.FormValue("synopsis")

	//link_amazon := r.FormValue("link_amazon")
	//link_indigo := r.FormValue("link_indigo")
	//link_pdf := r.FormValue("link_pdf")
	//link_epub := r.FormValue("link_epub")
	//link_handmade := r.FormValue("link_handmade")
	//link_text := r.FormValue("link_text")

	//verify title
	if !textNotEmpty(title) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify subtitle
	if !textNotEmpty(subtitle) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify author
	if !textNotEmpty(author) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify publish_year
	if !validYear(publish_year, publish_era) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify image
	if !verifyImage(handler) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify synopsis
	if !textNotEmpty(synopsis) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}

	//check book is already in the database
	rows, err := contentDB.Query("select title from books where title = ? and subtitle = ? and author = ?", title, subtitle, author)
	if err != nil {
		fmt.Printf("error checking if book exists: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	for rows.Next() {
		fmt.Printf("book already exists: %s\n", title)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	// Create the file in the file system
	dst, err := os.Create("../../public/assets/images/book_covers/" + handler.Filename)
	if err != nil {
		fmt.Printf("error creating the file: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	defer dst.Close()
	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Printf("error copying the file: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	//add book to database
	imagePath := "/assets/images/book_covers/" + handler.Filename
	_, err = contentDB.Exec("insert into books (title, subtitle, author, publish_year, image, synopsis) values (?, ?, ?, ?, ?, ?)", title, subtitle, author, publish_year, imagePath, synopsis)
	if err != nil {
		fmt.Printf("error adding book: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	fmt.Printf("successfully added book: %s\n", title)
	http.Redirect(w, r, "/addbook", http.StatusSeeOther)
}

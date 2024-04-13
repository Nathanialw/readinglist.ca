package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getAllBooksFromDB() []Book {
	books := []Book{}
	rows, err := contentDB.Query("SELECT title, subtitle, author, publish_year, image, synopsis FROM books")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Title, &book.Subtitle, &book.Author, &book.Publish_year, &book.Image, &book.Synopsis)
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)
	}

	return books
}

func addreadinglist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Books = getAllBooksFromDB()

	books := make([]Book, len(data.Books))
	for i, book := range data.Books {
		books[i] = Book{Title: book.Title, Subtitle: book.Subtitle, Author: book.Author, Image: book.Image, Synopsis: book.Synopsis}
	}
	// Encode the books slice to a JSON string
	booksJson, _ := json.Marshal(books)

	// Embed the JSON string in the data that will be passed to the template
	data.JsonBooks = string(booksJson)

	data.Categories, _ = Categories()
	generateHTML(w, data, "addreadinglist", "navbar", "footer", "addreadinglist")
}

func submitreadinglist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	fmt.Printf("successfully added reading list\n")
	http.Redirect(w, r, "/addreadinglist", http.StatusSeeOther)
}

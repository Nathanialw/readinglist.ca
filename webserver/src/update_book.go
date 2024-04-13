package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//get the book data from the DB
//display the book data in the form
//update the book data in the DB

func updatebook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Books = getAllBooksFromDB()

	books := make([]Book, len(data.Books))
	for i, book := range data.Books {
		books[i] = Book{Title: book.Title, Subtitle: book.Subtitle, Author: book.Author, Publish_year: book.Publish_year, Image: book.Image, Synopsis: book.Synopsis, Link_amazon: book.Link_amazon, Link_indigo: book.Link_indigo, Link_pdf: book.Link_pdf, Link_epub: book.Link_epub, Link_handmade: book.Link_handmade, Link_text: book.Link_text}
	}
	booksJson, _ := json.Marshal(books)
	data.JsonBooks = string(booksJson)

	generateHTML(w, data, "updatebook", "navbar", "footer", "updatebook")
}

func retievebook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	//	data.Book = GetBookFromDB(r, p.ByName("title"))

	generateHTML(w, data, "updatebook", "navbar", "footer", "updatebook")
}

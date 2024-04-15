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

func getAllBooksFromDBAllProperties() []Book {
	books := []Book{}
	rows, err := contentDB.Query("SELECT title, subtitle, author, publish_year, image, synopsis, link_amazon, link_indigo, link_pdf, link_epub, link_handmade, link_text FROM books")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Title, &book.Subtitle, &book.Author, &book.Publish_year, &book.Image, &book.Synopsis, &book.Link_amazon, &book.Link_indigo, &book.Link_pdf, &book.Link_epub, &book.Link_handmade, &book.Link_text)
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)
	}

	return books
}

func updatebook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Books = getAllBooksFromDBAllProperties()

	books := make([]Book, len(data.Books))
	for i, book := range data.Books {
		books[i] = Book{Title: book.Title, Subtitle: book.Subtitle, Author: book.Author, Publish_year: book.Publish_year, Image: book.Image, Synopsis: book.Synopsis, Link_amazon: book.Link_amazon, Link_indigo: book.Link_indigo, Link_pdf: book.Link_pdf, Link_epub: book.Link_epub, Link_handmade: book.Link_handmade, Link_text: book.Link_text}
	}
	booksJson, _ := json.Marshal(books)
	data.JsonBooks = string(booksJson)

	generateHTML(w, data, "updatebook", "navbar", "footer", "updatebook")
}

func submitupdatebook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	//check on the browser side first and give feedback to the user
	//then check on the server side
	//if this fails it should redirect to the addbook page

	var data UserSession
	data.LoggedIn = LoginStatus(r)

	if !data.LoggedIn {
		fmt.Println("not logged in\n")
		notfound(w, r, p)
		return
	}
	fmt.Printf("attempting to update book\n")
	if !VerifyAndInsertBook(w, r, submitDB) {
		fmt.Printf("adding book failed\n")
	}
	http.Redirect(w, r, "/updatebook", http.StatusSeeOther)
}

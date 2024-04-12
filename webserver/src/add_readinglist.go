package main

import (
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
	data.Books = getAllBooksFromDB()
	data.Categories, _ = Categories()
	generateHTML(w, data, "addreadinglist", "navbar", "footer", "addreadinglist")
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetReadingHistory(name string) (threads []Book, err error) {
	//get all the books from a readinglist
	//use the uid to get the books from the books table
	bookuids, err := contentDB.Query("select uid from readinghistory where user = ?", name)

	//get list of uids
	//use the uids to get the books

	var rows *sql.Rows

	for bookuids.Next() {
		lst := readinglistbooks{}
		bookuids.Scan(&lst.Bookuid)
		fmt.Printf("bookuid: %d\n", lst.Bookuid)
		rows, err = contentDB.Query("select title, subtitle, author, publish_year, image, synopsis, link_amazon, link_indigo, link_pdf, link_epub, link_handmade, link_text from books where uid = ?", lst.Bookuid)
		if err != nil {
			fmt.Printf("%s", err)
		}
		th := Book{}
		for rows.Next() {
			if err = rows.Scan(&th.Title, &th.Subtitle, &th.Author, &th.Publish_year, &th.Image, &th.Synopsis, &th.Link_amazon, &th.Link_indigo, &th.Link_pdf, &th.Link_epub, &th.Link_handmade, &th.Link_text); err != nil {
				fmt.Printf("%s", err)
				return
			}
			threads = append(threads, th)
			fmt.Printf("title: %s, author: %s\n", th.Title, th.Author)
		}
	}

	if rows != nil {
		rows.Close()
	}
	return
}

func ShowReadingHistory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.Books = getAllBooksFromDBAllProperties()

	generateHTML(w, data, "readinghistory", "readinghistory")
}

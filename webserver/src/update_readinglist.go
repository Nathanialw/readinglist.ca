package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getAllReadListsFromDBAllProperties() []ReadingList {
	readingLists := []ReadingList{}
	rows, err := contentDB.Query("SELECT name, category, image, chart, description FROM readinglists WHERE active = ?", "1")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		readingList := ReadingList{}
		err = rows.Scan(&readingList.Name, &readingList.Category, &readingList.Image, &readingList.Chart, &readingList.Description)
		if err != nil {
			fmt.Println(err)
		}
		readingLists = append(readingLists, readingList)
	}

	return readingLists
}

func updatereadinglist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//get the readinglist data from the DB
	//display the readinglist data in the form
	//update the readinglist data in the DB

	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.ReadingList = getAllReadListsFromDBAllProperties()

	data.Books = getAllBooksFromDB()

	books := make([]Book, len(data.Books))
	for i, book := range data.Books {
		books[i] = Book{Title: book.Title, Subtitle: book.Subtitle, Author: book.Author, Image: book.Image, Synopsis: book.Synopsis}
	}
	// Encode the books slice to a JSON string
	booksJson, _ := json.Marshal(books)
	data.JsonBooks = string(booksJson)

	readingLists := make([]ReadingList, len(data.ReadingList))
	for i, readingList := range data.ReadingList {
		readingLists[i] = ReadingList{Name: readingList.Name, Category: readingList.Category, Image: readingList.Image, Chart: readingList.Chart, Description: readingList.Description}
	}
	readingListsJson, _ := json.Marshal(readingLists)
	data.JsonReadingLists = string(readingListsJson)
	data.Categories, _ = Categories()

	generateHTML(w, data, "updatereadinglist", "navbar", "footer", "updatereadinglist")
}

func retrievereadinglist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//get the readinglist data from the DB
	//display the readinglist data in the form
	//update the readinglist data in the DB
}

func submitupdatereadinglist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func submitUpdatedReadingList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

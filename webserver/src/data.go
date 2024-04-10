package main

//create category
//create category chart
//create book info

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

type Category struct {
	Category    string
	Image       string
	Description string
}

type ReadingList struct {
	Name        string
	Category    string
	Image       string
	Chart       string
	Description string
}

type Book struct {
	Title         string
	Subtitle      string
	Author        string
	Publish_year  int
	Image         string
	Synopsis      string
	Link_amazon   string
	Link_indigo   string
	Link_pdf      string
	Link_epub     string
	Link_handmade string
	Link_text     string
}

type readinglistbooks struct {
	reading_list  string
	Bookuid       int
	reading_order int
}

type Reading struct {
	Reading_list ReadingList
	Books        []Book
}

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite", "../database/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
}

func Categories() (threads []Category, err error) {
	rows, err := db.Query("select name, image, description from categories where active = 1")

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	//fmt.Printf("adsa")
	for rows.Next() {
		th := Category{}
		if err = rows.Scan(&th.Category, &th.Image, &th.Description); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("category: %s, image: %s\n", th.Category, th.Image)
	}
	rows.Close()
	return
}

func ReadingLists(list string) (threads []ReadingList, err error) {
	rows, err := db.Query("select name, category, image, description from readinglists where category = ? and active = 1", list)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for rows.Next() {
		th := ReadingList{}
		if err = rows.Scan(&th.Name, &th.Category, &th.Image, &th.Description); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.Name, th.Category)
	}
	rows.Close()
	return
}

func Books(list string) (threads []Book, err error) {
	//get all the books from a readinglist
	//use the uid to get the books from the books table
	bookuids, err := db.Query("select bookuid from readinglistbooks where reading_list = ?", list)

	//get list of uids
	//use the uids to get the books

	var rows *sql.Rows

	for bookuids.Next() {
		lst := readinglistbooks{}
		bookuids.Scan(&lst.Bookuid)
		fmt.Printf("bookuid: %d\n", lst.Bookuid)
		rows, err = db.Query("select title, subtitle, author, publish_year, image, synopsis, link_amazon, link_indigo, link_pdf, link_epub, link_handmade, link_text from books where uid = ?", lst.Bookuid)
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

	rows.Close()
	return
}

func GetReadingList(name string) (threads ReadingList, err error) {
	rows, err := db.Query("select name, chart, description from readinglists where name = ?", name)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for rows.Next() {
		if err = rows.Scan(&threads.Name, &threads.Chart, &threads.Description); err != nil {
			fmt.Printf("%s", err)
			return
		}
		fmt.Printf("name: %s, image: %s\n", threads.Name, threads.Chart)
	}
	rows.Close()
	return
}

func GetCategory(name string) (threads Category, err error) {
	rows, err := db.Query("select name, image, description from categories where name = ? and active = 1", name)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for rows.Next() {
		if err = rows.Scan(&threads.Category, &threads.Image, &threads.Description); err != nil {
			fmt.Printf("%s", err)
			return
		}
		fmt.Printf("category: %s, image: %s\n", threads.Category, threads.Image)
	}
	rows.Close()
	return
}

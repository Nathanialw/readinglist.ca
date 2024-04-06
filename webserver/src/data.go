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
	rows, err := db.Query("select name, image, description from categories")

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

func ReadingLists() (threads []ReadingList, err error) {
	rows, err := db.Query("select name, category, image, description from categories")

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for rows.Next() {
		th := ReadingList{}
		if err = rows.Scan(&th.name, &th.category, &th.image, &th.description); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.name, th.category)
	}
	rows.Close()
	return
}

func Books() (threads []Book, err error) {
	rows, err := db.Query("select title, subtitle, author, publish_year, image, synopsis, link_amazon, link_indigo, link_pdf, link_epub, link_handmade from categories")

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for rows.Next() {
		th := Book{}
		if err = rows.Scan(&th.title, &th.subtitle, &th.author, &th.publish_year, &th.image, &th.synopsis, &th.link_amazon, &th.link_indigo, &th.link_pdf, &th.link_epub, &th.link_handmade); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("title: %s, author: %s\n", th.title, th.author)
	}
	rows.Close()
	return
}

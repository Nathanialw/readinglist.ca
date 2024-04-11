package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var userDB *sql.DB

type User struct {
	Username string
	Password string
}

func InsertIntoDB(r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Prepare SQL statement
	stmt, err := userDB.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return false
	}

	// Execute SQL statement
	_, err = stmt.Exec(username, password)
	return err == nil
}

func GetFromDB() bool {
	return true

}

func Authenticate(r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")
	rows, err := userDB.Query("select username, password from users where username = ? and password = ?", username, password)

	if err != nil {
		fmt.Printf("failed to authenticate: %s, %s, %s\n", username, password, err)
		rows.Close()
		return false
	}

	if rows.Next() {
		fmt.Printf("Successfully Authenticated: %s, %s\n", username, password)
		rows.Close()
		return true
	}
	fmt.Printf("failed to authenticate: %s, %s\n", username, password)
	rows.Close()
	return false
}

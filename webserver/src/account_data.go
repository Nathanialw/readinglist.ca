package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var userDB *sql.DB

type User struct {
	Username string
	Password string
}

var Username = "Nathan"

func FieldLength(r *http.Request) bool {
	//empty fields
	if len(r.FormValue("username")) == 0 || len(r.FormValue("email")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("repeat-password")) == 0 {
		fmt.Print("Emptyfield.\n")
		return false
	}
	//enforce formatting
	if len(r.FormValue("username")) <= 3 {
		fmt.Print("Username too short.\n")
		return false
	}
	if len(r.FormValue("username")) >= 15 {
		fmt.Print("Username too long.\n")
		return false
	}

	if len(r.FormValue("password")) <= 3 {
		fmt.Print("Password too short.\n")
		return false
	}
	if len(r.FormValue("password")) >= 15 {
		fmt.Print("Password too long.\n")
		return false
	}
	return true
}

func FieldLengthLogin(r *http.Request) bool {
	//empty fields
	if len(r.FormValue("username")) == 0 || len(r.FormValue("password")) == 0 {
		fmt.Print("Emptyfield.\n")
		return false
	}
	return true
}

func UserExists(r *http.Request) bool {
	rows, _ := userDB.Query("select username from users where username = ?", r.FormValue("username"))
	if rows.Next() {
		fmt.Print("User already exist.\n")
		rows.Close()
		return false
	}
	rows.Close()
	return true
}

func CheckPassword(r *http.Request) bool {
	if r.FormValue("password") != r.FormValue("repeat-password") {
		fmt.Print("Passwords do not match.\n")
		return false
	}
	return true
}

func CheckEmail(r *http.Request) bool {
	if !strings.Contains(r.FormValue("email"), "@") {
		fmt.Printf("%s, Not a valid email address.\n", r.FormValue("email"))
		return false
	}
	fmt.Printf("%s, This is a valid email address.\n", r.FormValue("email"))
	return true
}

func VerifyUserSignup(r *http.Request) bool {
	if !FieldLength(r) {
		return false
	}
	if !UserExists(r) {
		return false
	}
	if !CheckEmail(r) {
		return false
	}
	if !CheckPassword(r) {
		return false
	}

	return true
}

func VerifyUserLogin(r *http.Request) bool {
	if !FieldLengthLogin(r) {
		return false
	}
	if UserExists(r) {
		return false
	}

	return true
}

func InsertIntoDB(r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if !VerifyUserSignup(r) {
		return false
	}

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

	if !VerifyUserLogin(r) {
		return false
	}

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

func LoginStatus(r *http.Request) bool {
	cookie, err := r.Cookie("loggedin")
	if err != nil {
		// If there's an error, it means the cookie does not exist
		return false
	}

	// Check the value of the cookie
	if cookie.Value == "true" {
		return true
	} else {
		return false
	}
}

func SetCookies(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "loggedin",
		Value: "true",
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{
		Name:  "username",
		Value: r.FormValue("username"),
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
}

func ClearAllCookies(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "loggedin",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, cookie)
	cookie = &http.Cookie{
		Name:   "username",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, cookie)
}

func GetUsername(r *http.Request) string {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Printf("cookie: %s, error: %s\n", cookie.Value, err)
		return ""
	}
	fmt.Printf("cookie: %s\n", cookie.Value)
	return cookie.Value
}

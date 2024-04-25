package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetSavedReadingLists(user string) (threads []ReadingList, err error) {
	rows, err := contentDB.Query("select readinglist from savedreadinglists where user = ?", user)

	//if err != nil {
	//	fmt.Printf("%s", err)
	//	return
	//}
	//for rows.Next() {
	//	th := ReadingList{}
	//	if err = rows.Scan(&th.Name, &th.Category, &th.Image, &th.Description); err != nil {
	//		fmt.Printf("%s", err)
	//		return
	//	}
	//	threads = append(threads, th)
	//	fmt.Printf("name: %s, category: %s\n", th.Name, th.Category)
	//}
	if rows != nil {
		rows.Close()
	}
	return
}

func ShowSavedReadingLists(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	data.Admin = AdminStatus(r)
	data.ReadingList = getAllReadListsFromDBAllProperties()

	generateHTML(w, data, "savedreadinglists", "savedreadinglists")
}

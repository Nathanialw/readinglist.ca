package main

import "fmt"

func GetDailyReadingList() (th ReadingList, err error) {
	rows, err := contentDB.Query("select name, category, image, chart, description from readinglists where uid = ?", "16")

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for rows.Next() {
		th = ReadingList{}
		if err = rows.Scan(&th.Name, &th.Category, &th.Image, &th.Chart, &th.Description); err != nil {
			fmt.Printf("%s", err)
			return
		}
	}
	if rows != nil {
		rows.Close()
	}
	return
}

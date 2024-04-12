package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nfnt/resize"
)

func addbook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data UserSession
	data.LoggedIn = LoginStatus(r)
	generateHTML(w, data, "addbook", "navbar", "footer", "addbook")
}

func textNotEmpty(title string) bool {
	if title == "" {
		fmt.Println("title is empty")
		return false
	}
	return true
}

func validYear(year string, era string) bool {
	if year == "" || era == "" {
		fmt.Println("year is empty")
		return false
	}
	for _, c := range year {
		if c < '0' || c > '9' {
			fmt.Println("year is not a number")
			return false
		}
	}
	if era != "BC" && era != "AD" {
		fmt.Println("era is not BC or AD")
		return false
	}
	if len(year) > 4 {
		fmt.Println("year is more than 4 characters")
		return false
	}
	if year == "0" {
		fmt.Println("year is no year 0 BC or 0 AD")
		return false
	}
	if era == "AD" && year > "2024" {
		fmt.Println("the max year is 2024 AD")
		return false
	}
	return true
}

func verifyImage(handler *multipart.FileHeader) bool {
	if handler.Filename == "" {
		fmt.Println("image is empty")
		return false
	}
	// Check the file type
	fileType := handler.Header.Get("Content-Type")
	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		break
	default:
		fmt.Println("File is not an image")
		return false
	}
	if handler.Size > 4*1024*1024 {
		fmt.Printf("File is too large (max 4MB) %d\n", handler.Size)
		return false
	}
	//check if the file is an image
	if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
		fmt.Printf("file is not an image: %s\n", handler.Header.Get("Content-Type"))
		return false
	}
	//check if the file already exists
	if _, err := os.Stat("/assets/images/book_covers/" + handler.Filename); err == nil {
		//maybe append a number to the filename?
		fmt.Printf("file already exists: %s\n", handler.Filename)
		return false
	}

	return true
}

func submitbook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	//check on the browser side first and give feedback to the user
	//then check on the server side
	//if this fails it should redirect to the addbook page

	var data UserSession
	data.LoggedIn = LoginStatus(r)

	if !data.LoggedIn {
		fmt.Println("not logged in")
		notfound(w, r, p)
		return
	}

	//verify book data is formatted correctly
	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	author := r.FormValue("author")
	publish_year := r.FormValue("publish_year")
	publish_era := r.FormValue("publish_era")
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	// Retrieve the file from form data
	file, handler, err := r.FormFile("image") // "image" is the name of the file input field
	if err != nil {
		fmt.Printf("error retrieving the file: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	synopsis := r.FormValue("synopsis")

	//link_amazon := r.FormValue("link_amazon")
	//link_indigo := r.FormValue("link_indigo")
	//link_pdf := r.FormValue("link_pdf")
	//link_epub := r.FormValue("link_epub")
	//link_handmade := r.FormValue("link_handmade")
	//link_text := r.FormValue("link_text")

	//verify title
	if !textNotEmpty(title) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify subtitle
	if !textNotEmpty(subtitle) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify author
	if !textNotEmpty(author) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify publish_year
	if !validYear(publish_year, publish_era) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify image
	if !verifyImage(handler) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}
	//verify synopsis
	if !textNotEmpty(synopsis) {
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
	}

	//check book is already in the database
	rows, err := contentDB.Query("select title from books where title = ? and subtitle = ? and author = ?", title, subtitle, author)
	if err != nil {
		fmt.Printf("error checking if book exists: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	for rows.Next() {
		fmt.Printf("book already exists: %s\n", title)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	// Create the file in the file system
	systemPath := "../../public/assets/images/book_covers/" + handler.Filename
	dst, err := os.Create(systemPath)
	if err != nil {
		fmt.Printf("error creating the file: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}
	defer dst.Close()
	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Printf("error copying the file: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	OSFile, _ := os.Open(systemPath)
	defer file.Close()

	//resize the image
	img, _, err := image.Decode(OSFile)
	if err != nil {
		fmt.Printf("error decoding the image: %s, %s\n", err, systemPath)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	extension := filepath.Ext(systemPath)
	switch extension {
	case ".jpg", ".jpeg":
		fmt.Printf("jpeg\n")
		OSFile.Seek(0, 0) // Reset the reader to the start of the file
		img, err = jpeg.Decode(OSFile)
		//resize the image to 400x400
		m := resize.Resize(400, 0, img, resize.Lanczos3)
		systemPath = "../../public/assets/images/book_covers/400_" + handler.Filename
		out, _ := os.Create(systemPath)
		defer out.Close()
		// Write the new image to the new file
		jpeg.Encode(out, m, nil)

		m = resize.Resize(100, 0, img, resize.Lanczos3)
		systemPath = "../../public/assets/images/book_covers/100_" + handler.Filename
		out, _ = os.Create(systemPath)
		jpeg.Encode(out, m, nil)
	case ".png":
		fmt.Printf("png\n")
		OSFile.Seek(0, 0) // Reset the reader to the start of the file
		img, err = png.Decode(OSFile)
		//resize the image to 400x400
		m := resize.Resize(400, 0, img, resize.Lanczos3)
		systemPath = "../../public/assets/images/book_covers/400_" + handler.Filename
		out, _ := os.Create(systemPath)
		defer out.Close()
		// Write the new image to the new file
		png.Encode(out, m)

		m = resize.Resize(100, 0, img, resize.Lanczos3)
		systemPath = "../../public/assets/images/book_covers/100_" + handler.Filename
		out, _ = os.Create(systemPath)
		png.Encode(out, m)
	case ".gif":
		fmt.Printf("unsupported image format: %s\n", extension)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
		//OSFile.Seek(0, 0) // Reset the reader to the start of the file
		//img, err = gif.Decode(OSFile)
		////resize the image to 400x400
		//m := resize.Resize(400, 0, img, resize.Lanczos3)
		//systemPath = "../../public/assets/images/book_covers/400_" + handler.Filename
		//out, _ := os.Create(systemPath)
		//defer out.Close()
		//// Write the new image to the new file
		//gif.Encode(out, m, nil)
		//
		//m = resize.Resize(100, 0, img, resize.Lanczos3)
		//systemPath = "../../public/assets/images/book_covers/100_" + handler.Filename
		//out, _ = os.Create(systemPath)
		//gif.Encode(out, m, nil)
	default:
		fmt.Printf("unsupported image format: %s\n", extension)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	//add book to database
	imagePath := "/assets/images/book_covers/" + handler.Filename
	imagePath100 := "/assets/images/book_covers/100_" + handler.Filename
	imagePath400 := "/assets/images/book_covers/400_" + handler.Filename
	_, err = contentDB.Exec("insert into books (title, subtitle, author, publish_year, image, image_100, image_400, synopsis) values (?, ?, ?, ?, ?, ?, ?, ?)", title, subtitle, author, publish_year, imagePath, imagePath100, imagePath400, synopsis)
	if err != nil {
		fmt.Printf("error adding book: %s\n", err)
		http.Redirect(w, r, "/addbook", http.StatusSeeOther)
		return
	}

	fmt.Printf("successfully added book: %s\n", title)
	http.Redirect(w, r, "/addbook", http.StatusSeeOther)
}

function  IsSet(value) {
    let text = "placeholder=\"Not Set\""
    if (value != "" && value != "NA") {
        text = "value=\"" + value + "\""
    }
    return text
}

function  IsSetRawText(value) {
    if (value == "" || value == "NA") {
        return ""
    }
    return value
}

function SetHTML (selectedBook) {
    document.getElementById("book-form").innerHTML = `
    <form class="form" action="/submitupdatebook" method="post" enctype="multipart/form-data">
    <label class="form-label" for="title">Title</label>
    <input class="form-input" type="text" id="title" name="title" ${IsSet(selectedBook.Title)} required>
    <label class="form-label" for="title">Subtitle</label>
    <input class="form-input" type="text" id="subtitle" name="subtitle" ${IsSet(selectedBook.Subtitle)}>
    <label class="form-label" for="author">Author</label>
    <input class="form-input" type="text" id="author" name="author" ${IsSet(selectedBook.Author)}>
    <span class="form-adjacent" style="margin: 1rem auto;">
        <label class="form-label form-adjacent" for="year">Year Published</label>
        <input class="form-input form-adjacent" style="width: 150px;" type="number" min="1" max="6000" id="year" name="publish_year" ${IsSet(selectedBook.Publish_year)} required>
        <select class="form-input form-adjacent" style="text-transform: capitalize; width: 125px; font-size: 0.75rem" id="publish_era" name="publish_era" required>
            <option value="AD">AD</option>
            <option value="BC">BC</option>
        </select>
    </span>
    <label class="form-label" for="image">Book Cover:</label>
    <span class="form-adjacent" style="margin: 1rem auto;">
        <img class="form-image" id="image-preview" src="${selectedBook.Image}" alt="Book Cover Preview">
        <input class="form-input form-row" type="file" id="image" name="image" ${IsSet(selectedBook.Image)}>
    </span>
    <label class="form-label" for="synopsis">Synopsis</label>
    <textarea class="form-input form-desc" id="synopsis" name="synopsis" ${SetTextAreaHieght(selectedBook)} style="resize: none;" ${IsSet(selectedBook.Synopsis)} required>${IsSetRawText(selectedBook.Synopsis)}</textarea>
    <label class="form-label" for="link_amazon">Amazon Link</label>
    <input class="form-input" type="url" id="link_amazon" name="link_amazon" ${IsSet(selectedBook.Link_amazon)}>
    <label class="form-label" for="link_indigo">Indigo Link</label>
    <input class="form-input" type="url" id="link_indigo" name="link_indigo" ${IsSet(selectedBook.Link_indigo)}>
    <label class="form-label" for="link_pdf">PDF Link</label>
    <input class="form-input" type="url" id="link_pdf" name="link_pdf" ${IsSet(selectedBook.Link_pdf)}>
    <label class="form-label" for="link_epub">EPUB Link</label>
    <input class="form-input" type="url" id="link_epub" name="link_epub" ${IsSet(selectedBook.Link_epub)}">
    <label class="form-label" for="link_text">Text Link</label>
    <input class="form-input" type="url" id="link_text" name="link_text" ${IsSet(selectedBook.Link_text)}">
    <input class="btn btn-center btn-submit" type="submit" value="Update Book">
    </form>
    `
    document.getElementById("image").addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var imageUrl = URL.createObjectURL(selectedFile);
        document.getElementById("image-preview").src = imageUrl;
    });
}




function SetTextAreaHieght (selectedBook) {
    let description = IsSetRawText(selectedBook.Synopsis);
    let numChars = description.length;
    let charsPerLine = 50; // adjust this value based on the width of the textarea and the size of the font
    let numLines = Math.ceil(numChars / charsPerLine);

    return "rows=\"" + numLines + "\""
}

//I need to comme up with a way to dynamically add more lines to the height textarea when the user types more text without redoing the whol form html

function SetTextAreaHieght (selectedBook) {
    let description = IsSetRawText(selectedBook.Synopsis);
    let numChars = description.length;
    let charsPerLine = 50; // adjust this value based on the width of the textarea and the size of the font
    let numLines = Math.ceil(numChars / charsPerLine);

    return "rows=\"" + numLines + "\""
}



// function UpdateTextAreaHeight (selectedBook) {
//     document.getElementById("description").innerHTML = `
//     <textarea class="form-input form-desc" id="description" name="description" style="resize: none;" ${IsSet(selectedBook.Synopsis)} required>${IsSetRawText(selectedBook.Synopsis)}</textarea>
//     `

//     let textarea = document.getElementById("description");
//     textarea.addEventListener("input", function() {
//         let numLines = (this.value.match(/\n/g) || []).length + 1;
//         this.rows = numLines;
//     });
// }



//replace the html in the select options with the sorted array
function displayBooks(bookList) {
    //save the current value of the select
    var selectedOption = document.getElementById(bookList).value;
    
    var bookList = document.getElementById(bookList);
    while (bookList.firstChild) {
        bookList.removeChild(bookList.firstChild);
    }
    bookList.innerHTML = "";
    for (var i = 0; i < books.length; i++) {
        var option = document.createElement("option");
        option.value = books[i].Title;
        let text = books[i].Title
        if (books[i].Subtitle != "" && books[i].Subtitle != "NA") {
            text += ": " + books[i].Subtitle
        }
        text += ", by: " + books[i].Author
        option.text = text;
        bookList.add(option);
    }
    //after sorting, set the value of the select to the saved value
    if (selectedOption == "") {
        selectedOption = books[0].Title
        // selectedOption = "Select a book..."
        console.log(selectedOption)
    }
    document.getElementById(bookList.id).value = selectedOption;
    let selectedBook = books.find(book => book.Title === selectedOption)

    SetHTML(selectedBook)
}

displayBooks("bookList")



//get the books from the server
document.getElementById("bookList").addEventListener("change", function () {
    selectedOption = document.getElementById(bookList.id).value;
    let selectedBook = books.find(book => book.Title === selectedOption)
    SetHTML(selectedBook)
})







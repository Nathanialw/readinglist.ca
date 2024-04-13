













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
    }
    document.getElementById(bookList.id).value = selectedOption;
}

displayBooks("bookList")
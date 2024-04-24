'use strict'

function  IsSet(value) {
    return sanitizeHTML(value)
}

function IsSetRawText(value) {
    if (value == "" || value == "NA") {
        value = ""
    }
    return sanitizeHTML(value)
}

function sanitizeHTML(text) {
    let div = document.createElement('div');
    div.innerText = text;
    return div.innerHTML;
}

function sanitizeImageUrl(url) {
    return isValidImageUrl(url) ? url : '';
}

function isValidImageUrl(url) {
    try {
        new URL(url, window.location.href);
    } catch (_) {
        return false;  
    }
    return url.endsWith('.png') || url.endsWith('.jpg') || url.endsWith('.jpeg') || url.endsWith('.gif');
}

function SetHTML(selectedReadingList) {
    let name = document.getElementById("name")
    name.setAttribute("value", IsSet(selectedReadingList.Name))
    let imagepreview = document.getElementById("image-preview")
    imagepreview.setAttribute("src", sanitizeImageUrl(selectedReadingList.Chart))
    let description = document.getElementById("description")
    description.innerText = IsSet(selectedReadingList.Description)
    description.setAttribute("rows", SetTextAreaHieght(selectedReadingList))

    document.getElementById("image").addEventListener("change", function (event) {
        var selectedFile = event.target.files[0];
        var imageUrl = URL.createObjectURL(selectedFile);
        document.setAttribute("src", imageUrl)
    });
    
    document.getElementById("description").addEventListener("input", () => {
        let numChars = description.value.length;
        let charsPerLine = 70; // adjust this value based on the width of the textarea and the size of the font
        let numLines = Math.ceil(numChars / charsPerLine);
        if (description.value.includes('\n')) {
            numLines += description.value.split('\n').length - 1;
        }
        numLines < 4 ? numLines = 4 : numLines
        description.setAttribute("rows", numLines)
    });
}

function SetTextAreaHieght(selectedReadingList) {
    let description = IsSetRawText(selectedReadingList.Description);
    let numChars = description.length;
    let charsPerLine = 70; // adjust this value based on the width of the textarea and the size of the font
    let numLines = Math.ceil(numChars / charsPerLine);

    numLines < 4 ? numLines = 4 : numLines
    return numLines
}

//replace the html in the select options with the sorted array
function displayReadingList(readingList) {
    //save the current value of the select
    var selectedOption = document.getElementById(readingList).value;
    
    var readingList = document.getElementById(readingList);
    while (readingList.firstChild) {
        readingList.removeChild(readingList.firstChild);
    }
    readingList.innerHTML = "";
    for (var i = 0; i < readinglists.length; i++) {
        var option = document.createElement("option");
        option.value = readinglists[i].Name;
        let text = readinglists[i].Name
        // if (readinglists[i].Subtitle != "" && readinglists[i].Subtitle != "NA") {
        //     text += ": " + readinglists[i].Subtitle
        // }
        text += ", " + readinglists[i].Category
        option.text = text;
        readingList.add(option);
    }
    //after sorting, set the value of the select to the saved value
    if (selectedOption == "") {
        selectedOption = readinglists[0].Name
        // selectedOption = "Select a book..."
    }
    document.getElementById(readingList.id).value = selectedOption;
    let selectedReadingList = readinglists.find(readinglist => readinglist.Name === selectedOption)
    SetHTML(selectedReadingList)
}

displayReadingList("readingLists")

//get the books from the server
document.getElementById("readingLists").addEventListener("change", function () {
    var selectedOption = document.getElementById("readingLists").value;
    let selectedReadingList = readinglists.find(readinglist => readinglist.Name === selectedOption)
    SetHTML(selectedReadingList)
})




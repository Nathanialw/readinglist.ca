//   ////////////////
const book_name = {
  title: "",
  synopsis: "",
  subtitle: "",
  author: "",
  image: "",
};

document.getElementById("author").innerHTML = book_name.author;

document.getElementById("title").innerHTML = book_name.title;

document.getElementById("subtitle").innerHTML = book_name.subtitle;

document.getElementById("synopsis").innerHTML = book_name.synopsis;

document.getElementById("image").innerHTML =
  "<img src='" + book_name.image + "' class='img' alt=''>";

///////////////////////////////////

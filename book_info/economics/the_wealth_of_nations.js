const the_wealth_of_nations = {
  title: "the wealth of nations",
  subtitle: "An Inquiry Into the Nature and Causes of the Wealth of Nations",
  author: "adam smith",
  synopsis:
    "Adam Smith's The Wealth of Nations was recognized as a landmark of human thought upon its publication in 1776. As the first scientific argument for the principles of political economy, it is the point of departure for all subsequent economic thought. Smith's theories of capital accumulation, growth, and secular change, among others, continue to be influential in modern economics.This reprint of Edwin Cannan's definitive 1904 edition of The Wealth of Nations includes Cannan's famous introduction, notes, and a full index, as well as a new preface written especially for this edition by the distinguished economist George J. Stigler. Mr. Stigler's preface will be of value for anyone wishing to see the contemporary relevance of Adam Smith's thought.",
  image:
    "/categories/acedemic/index/economics/images/the-wealth-of-nations.jpg",
};

document.getElementById("the_wealth_of_nations_author").innerHTML =
  the_wealth_of_nations.author;

document.getElementById("the_wealth_of_nations_title").innerHTML =
  the_wealth_of_nations.title;

document.getElementById("the_wealth_of_nations_subtitle").innerHTML =
  the_wealth_of_nations.subtitle;

document.getElementById("the_wealth_of_nations_synopsis").innerHTML =
  the_wealth_of_nations.synopsis;

document.getElementById("the_wealth_of_nations_image").innerHTML =
  "<img src='" + the_wealth_of_nations.image + "' class='img' alt=''>";

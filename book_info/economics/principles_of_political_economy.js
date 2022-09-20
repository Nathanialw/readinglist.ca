const principles_of_political_economy = {
  title: "principles of political economy",
  subtitle: "with some of their applications to Social Philosophy",
  author: "john stuart mill",
  synopsis:
    "The Principles of Political Economy with some of their applications to Social Philosophy by John Stuart Mill. Principles of Political Economy (1848) by John Stuart Mill was one of the most important economics or political economy textbooks of the mid-nineteenth century. It was revised until its seventh edition in 1871, shortly before Mill's death in 1873, and republished in numerous other editions. In every department of human affairs, Practice long precedes Science systematic enquiry into the modes of action of the powers of nature, is the tardy product of a long course of efforts to use those powers for practical ends. The conception, accordingly, of Political Economy as a branch of science is extremely modern; but the subject with which its enquiries are conversant has in all ages necessarily constituted one of the chief practical interests of mankind, and, in some, a most unduly engrossing one. That subject is Wealth. Writers on Political Economy profess to teach, or to investigate, the nature of Wealth, and the laws of its production and distribution: including, directly or remotely, the operation of all the causes by which the condition of mankind, or of any society of human beings, in respect to this universal object of human desire, is made prosperous or the reverse. Not that any treatise on Political Economy can discuss or even enumerate all these causes; but it undertakes to set forth as much as is known of the laws and principles according to which they operate.",
  image:
    "/categories/acedemic/index/economics/images/principles of-political-economy.webp",
};

document.getElementById("principles_of_political_economy_author").innerHTML =
  principles_of_political_economy.author;

document.getElementById("principles_of_political_economy_title").innerHTML =
  principles_of_political_economy.title;

document.getElementById("principles_of_political_economy_subtitle").innerHTML =
  principles_of_political_economy.subtitle;

document.getElementById("principles_of_political_economy_synopsis").innerHTML =
  principles_of_political_economy.synopsis;

document.getElementById("principles_of_political_economy_image").innerHTML =
  "<img src='" +
  principles_of_political_economy.image +
  "' class='img' alt=''>";

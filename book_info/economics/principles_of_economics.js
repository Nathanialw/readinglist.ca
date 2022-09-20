const principles_of_economics = {
  title: "principle of economics",
  subtitle: "",
  author: "carl menger",
  synopsis:
    "In the beginning, there was Menger. It was his Principles of Economics that reformulated — and really rescued — economic science from the theoretical errors of the old classical school.Menger set out to elucidate the precise nature of economic value, and to root economics firmly in the real-world actions of individual human beings. In Principles of Economics, he advances the theory that the marginal utility of goods is the source of their value, rather than the labor inputs that went into making them. The implication of this theory is that the individual mind is the source of economic value — a point that touched off the marginalist revolution and started a departure from the flawed classical view of economics.For this reason, Carl Menger (1840–1921) is considered to be the founder of the Austrian School of economics. Principles of Economics is the book that Ludwig von Mises said turned him into a real economist. What's striking is how, nearly a century and a half later, the book still retains the incredible power of both its prose and its relentless logic.The Mises Institute's new edition features a new foreword by Peter G. Klein, which summarizes Menger's contribution and places him in the history of ideas. Klein also explains Menger's continued relevance in present times. F.A. Hayek contributes the introduction.Economics students still say that it is the best introduction to economic logic ever written. The book also deserves the status of a seminal contribution to science in general. Truly, no one can claim to be well read in economics without having mastered Menger's argument.To search for Mises Institute titles, enter a keyword and LvMI (short for Ludwig von Mises Institute); e.g., Depression LvMI",
  image:
    "/categories/acedemic/index/economics/images/principles-of-economics.jpg",
};

document.getElementById("principles_of_economics_title").innerHTML =
  principles_of_economics.title;

document.getElementById("principles_of_economics_subtitle").innerHTML =
  principles_of_economics.subtitle;

document.getElementById("principles_of_economics_author").innerHTML =
  principles_of_economics.author;

document.getElementById("principles_of_economics_synopsis").innerHTML =
  principles_of_economics.synopsis;

document.getElementById("principles_of_economics_image").innerHTML =
  "<img src='" + principles_of_economics.image + "' class='img' alt=''>";

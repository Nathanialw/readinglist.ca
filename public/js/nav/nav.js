{
  const nav =
    "<div class='navbar'>\
      <a href='/index.html'><h4 class='logo'>logo</h4></a\
      ><button type='button' class='bars' id='bars'>\
        <i class='fa fa-bars'></i>\
      </button>\
    </div>\
    <nav>\
      <ul class='nav-links' id='nav-links'>\
        <li><a href='/index.html'>home</a></li>\
        <li><a href='/about.html'>about</a></li>\
        <li><a href='/contact.html'>contact</a></li>\
      </ul>\
    </nav>";

  document.getElementById("nav-bar").insertAdjacentHTML("beforeend", nav);

  const getElement = (selector) => {
    const el = document.querySelector(selector);
    if (el) return el;
    throw new Error(`Please check your classes : ${selector} does not exist`);
  };

  const navToggle = getElement(".bars");
  const links = getElement(".nav-links");

  navToggle.addEventListener("click", function () {
    console.log("henlo");
    links.classList.toggle("show-links");
  });

  /*
  
  <link rel="stylesheet" href="/nav/style.css" />
  <nav id="nav-bar"></nav>
  <script src="/nav/nav.js"></script>

  */
}

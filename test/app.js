{
  let test = document.getElementById("test");
  let node = document.createElement("li");
  let list = document.getElementById("nav-links");

  test.onmouseover = function () {
    console.log(node);
    node.appendChild(document.createTextNode("new"));

    list.appendChild(node);
  
    test.onclick = function () {
      list.removeChild(node);
    };
  };
}

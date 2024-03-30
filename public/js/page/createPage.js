{

    const cards = document.querySelector('.cardss')

    function createPage(books) {
        books.map((book) => {    
            cards.innerHTML = books.map(({title: title, subtitle: subtitle, author: author, synopsis: synopsis, image: image, link_pdf: pdf, link_epub: epub, link_amazon: amazon, link_indigo: indigo, link_handmade: handmade}) => {
            return `
            <li>
                <div class="card">
                
                <div class="book-img">
                    <img src=${image} class='img' alt=''>
                </div>
            
                <div class="book-info">
                    <h3 class="book-title"">${title}</h3>
                    <h4 class="book-subtitle">${subtitle}</h4>
                    <div class="title-underline"></div>
                    <h4 class="book-author">${author}</h4>
                    <ul class="buy-links">    
                    <li><a class='fa fa-file-pdf' href=${pdf} target='_blank'></a></li>
                    <li><a class='fa fa-book-open' href=${epub} target='_blank' ></a></li>
                    <li><a class='fab fa-amazon' href=${amazon} target='_blank' ></a></li>
                    <li><a class='' href='${indigo}' target='_blank'>indigo</a></li>
                    <li><a class='fa fa-book' href=${handmade} target='_blank'></a></li>
                    </ul>
                </div>  
            
                </div>
            
                <p class="book-synopsis">${synopsis}</p>
            </li>`
            }).join('')
        })
    }
}
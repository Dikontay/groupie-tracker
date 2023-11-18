let currentFocus = -1

const searchInput = document.getElementById('search-input')
const suggestionsElement = document.getElementById('suggestions')

searchInput.addEventListener('input', function(event){
    currentFocus=-1
    const query = event.target.value

    if (query.length > 0) {
        fetch('http://localhost:4000/search?query=${encodeURIComponent(query)}')
        .then(response => response.json())
        .then(suggesions => {
            suggestionsElement.innerHTML = '';
            suggestions.forEach((suggestion, index) => {
                const div = document.createElement('div')
                div.textContent = suggestion;
                div.classList.add('suggestions-item);
                
            });
        })
    }
})
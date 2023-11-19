const searchInput = document.getElementById('search-input');
const suggestionsElement = document.getElementById('suggestions');

searchInput.addEventListener('input', function(event) {
    currentFocus = -1; // Reset the current focus
    const query = event.target.value;

    if (query.length > 0) {
        fetch(`http://localhost:4000/search?query=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(suggestions => {
                suggestionsElement.innerHTML = '';
                suggestions.forEach((suggestion, index) => {
                    const div = document.createElement('div');
                    div.textContent = suggestion;
                    div.classList.add('suggestion-item');
                    div.onclick = function() {
                        searchInput.value = suggestion;
                        suggestionsElement.innerHTML = '';
                    };
                    suggestionsElement.appendChild(div);
                });
            })
            .catch(error => console.error('Error:', error));
    } else {
        suggestionsElement.innerHTML = '';
    }
});

searchInput.addEventListener('keydown', function(e) {
    let items = suggestionsElement.getElementsByClassName('suggestion-item');
    if (e.keyCode == 40) { // Down key
        currentFocus++;
        addActive(items);
    } else if (e.keyCode == 38) { // Up key
        currentFocus--;
        addActive(items);
    } else if (e.keyCode == 13) { // Enter key
        e.preventDefault();
        if (currentFocus > -1) {
            if (items) items[currentFocus].click();
        }
    }
});

function addActive(items) {
    if (!items) return false;
    removeActive(items);
    if (currentFocus >= items.length) currentFocus = 0;
    if (currentFocus < 0) currentFocus = (items.length - 1);
    items[currentFocus].classList.add('suggestion-active');
}

function removeActive(items) {
    for (let i = 0; i < items.length; i++) {
        items[i].classList.remove('suggestion-active');
    }
}

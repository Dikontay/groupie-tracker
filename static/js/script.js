document.getElementById('search-input').addEventListener('input', function(event) {
    const query = event.target.value;
    if (query.length > 0) {
        fetch(`http://localhost:8000/search?query=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(suggestions => {
                const suggestionsElement = document.getElementById('suggestions');
                suggestionsElement.innerHTML = ''; // Clear previous suggestions
                suggestionsElement.style.display = 'block'; // Show suggestions list
                suggestions.forEach(suggestion => {
                    const li = document.createElement('li');
                    li.textContent = suggestion;
                    li.classList.add('list-group-item');
                    li.addEventListener('click', () => {
                        document.getElementById('search-input').value = suggestion;
                        suggestionsElement.innerHTML = ''; // Clear suggestions
                        suggestionsElement.style.display = 'none'; // Hide suggestions list
                    });
                    suggestionsElement.appendChild(li);
                });
            })
            .catch(error => console.error(error));
    } else {
        document.getElementById('suggestions').style.display = 'none'; // Hide suggestions list if query is empty
    }
});
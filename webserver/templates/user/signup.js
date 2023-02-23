const form = document.querySelector('.signup form');

form.addEventListener('submit', (event) => {
    event.preventDefault();

    // Extract form data
    const formData = new FormData(form);
    const data = {
        name: formData.get('name'),
        email: formData.get('email'),
        password: formData.get('psword')
    };

    // Send POST request to the server
    fetch('http://localhost/v1/user/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (response.ok) {
                // Handle successful response
                return response.json();
            } else {
                // Handle error response
                throw new Error('Error signing up user');
            }
        })
        .then(data => {
            // Log token
            console.log('Token:', data.token);
        })
        .catch(error => {
            // Handle error
            console.error(error);
        });
});
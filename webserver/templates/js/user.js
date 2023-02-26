const form = document.querySelector('.signup-form');
const loginFormEl = document.querySelector('.login-form');

const url = 'http://todotask.server'

function handleAuthentication(type) {
    if (type === 'signUp') {
        // Extract form data
        const formData = new FormData(form);
        const data = {
            name: formData.get('name'),
            email: formData.get('email'),
            password: formData.get('psword')
        };

        // Send POST request to the server
        fetch(`${url}/v1/user/signup`, {
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
                localStorage.setItem('Token', data.token);
                location.href = 'taskmanager.html';
            })
            .catch(error => {
                // Handle error
                console.error(error);
            });
    }
    else if(type === 'login'){
        // Extract form data
        const formData = new FormData(loginFormEl);
        const data = {
            email: formData.get('email'),
            password: formData.get('password')
        };

        // Send POST request to the server
        fetch(`${url}/v1/user/signin`, {
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
                localStorage.setItem('Token', data.token);
                location.href = 'taskmanager.html';
            })
            .catch(error => {
                // Handle error
                console.error(error);
            });
    }
}

form.addEventListener('submit', (event) => {
    event.preventDefault();
    handleAuthentication('signUp');
});

loginFormEl.addEventListener('submit', (event) => {
    event.preventDefault();
    handleAuthentication('login')
})
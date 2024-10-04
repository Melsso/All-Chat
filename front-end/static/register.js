document.addEventListener('DOMContentLoaded', function() {
    const registerForm = document.getElementById('register-form');

    if (registerForm) {
        registerForm.addEventListener('submit', function(event) {
            event.preventDefault();

            const firstName = document.getElementById('first-name').value;
            const lastName = document.getElementById('last-name').value;
            const dob = document.getElementById('dob').value; 
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            const gender = document.getElementById('gender').value;

            if (!firstName || !lastName || !dob || !email || !password || !gender) {
                alert('Please fill in all fields');
                return;
            }

            fetch('http://localhost:8000/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ 
                    first_name: firstName,
                    last_name: lastName,
                    date_of_birth: dob,
                    email: email,
                    password: password,
                    gender: gender
                }),
            })
            .then(response => {
                if (!response.ok) {
                    return response.json().then(err => { throw err; });
                }
                return response.json();
            })
            .then(data => {
                console.log('Registration successful:', data);
                window.location.href = '/static/login.html';
            })
            .catch(error => {
                console.error('Error:', error);
                alert('Registration failed. Please try again.');
            });
        });
    } 
});

// login.js

function login() {
    // Obtener los valores de usuario y contraseña
    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    // Verificar que ambos campos estén completos
    if (!username || !password) {
        alert("Por favor, ingresa nombre de usuario y contraseña.");
        return;
    }

    // Construir el objeto JSON con la información requerida
    var loginData = {
        "username": username,
        "password": password
    };

    // Realizar la solicitud HTTP POST
    fetch('http://localhost:8090/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(loginData)
    })
    .then(response => response.json())
    .then(data => {
        // Verificar si la solicitud fue exitosa
        if (data.token) {
            // Guardar el token en el localStorage
            localStorage.setItem('token', data.token);
            localStorage.setItem('userId', data.userId);

            // Redirigir a la página "vender.html"
            window.location.href = 'vender.html';
        } else {
            alert("Credenciales inválidas. Por favor, inténtalo nuevamente.");
        }
    })
    .catch(error => {
        console.error('Error al iniciar sesión:', error);
        alert("Error al iniciar sesión. Por favor, inténtalo nuevamente.");
    });
}
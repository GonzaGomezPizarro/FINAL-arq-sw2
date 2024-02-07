//vender.js

// Verificar si hay un token almacenado
const token = localStorage.getItem('token');
console.log('Valor del token:', token);

if (!token) {
    // Si no hay token, redirigir a la página de inicio de sesión
    console.log('No hay token, redirigiendo a login.html');
    window.location.href = 'login.html';
}

function verPublicaciones() {
    // Redirigir a la página de publicaciones
    window.location.href = 'publicaciones.html';
}

function nuevaPublicacion() {
    // Redirigir a la página de nueva publicación   
    window.location.href = 'nuevapublicacion.html';
}

function volver(){
    window.location.href = 'index.html';
}
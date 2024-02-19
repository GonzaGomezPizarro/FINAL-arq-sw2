// publicacion.js

// Recuperar el resultado almacenado en el almacenamiento local
var publicacion = JSON.parse(localStorage.getItem('publicacion'));

console.log(publicacion)

// Obtener el elemento principal donde se mostrarán los detalles
var detalleContainer = document.getElementById('detalleContainer');

// Mostrar las fotos de la propiedad
var fotosContainer = document.createElement('div');
fotosContainer.className = 'imagenes';

publicacion.photos.forEach(function(fotoBase64) {
    // Crear una imagen y configurar su fuente como una URL de imagen basada en datos base64
    var imagen = document.createElement('img');
    imagen.src = 'data:image/jpeg;base64,' + fotoBase64;
    fotosContainer.appendChild(imagen);
});

detalleContainer.appendChild(fotosContainer);

// Mostrar los detalles de la propiedad
var campos = ['title', 'description', 'country', 'state', 'city', 'address', 'price', 'bedrooms', 'bathrooms', 'mts2', 'userId'];

campos.forEach(function(campo) {
    var etiqueta = document.createElement('p');
    etiqueta.textContent = `${campo.charAt(0).toUpperCase() + campo.slice(1)}: ${publicacion[campo]}`;
    detalleContainer.appendChild(etiqueta);
});


// Limpiar el resultado almacenado en el almacenamiento local después de mostrar los detalles
//localStorage.removeItem('publicacion');

function volver(){
    window.location.href = 'publicaciones.html';
}


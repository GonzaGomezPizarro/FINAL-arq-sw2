// publicaciones.js

// Obtener userId almacenado en localStorage
const userId = localStorage.getItem('userId');
console.log('userId:', userId);

// Verificar si hay un userId almacenado
if (!userId) {
    // Si no hay userId, redirigir a la página de inicio de sesión
    console.log('No hay userId, redirigiendo a login.html');
    window.location.href = 'login.html';
}

// Construir la URL para la consulta
const apiUrl = `http://localhost:8000/items/${userId}`;

// Realizar una solicitud HTTP GET utilizando Fetch
fetch(apiUrl)
    .then(response => response.json())
    .then(data => {
        console.log('Publicaciones obtenidas:', data);
        if (data == null){
              // Si no hay publicaciones, mostrar un mensaje indicando que no hay nada publicado
              var mensajeNoPublicaciones = document.createElement('p');
              mensajeNoPublicaciones.textContent = 'No hay publicaciones disponibles en este momento.';
              document.getElementById('publicacionesContainer').appendChild(mensajeNoPublicaciones);
              return;
        }
        mostrarPublicaciones(data);
    })
    .catch(error => {
        console.error('Error al realizar la solicitud:', error);
    });

    // ...

function mostrarPublicaciones(publicaciones) {
    // Obtener el elemento principal donde se mostrarán las publicaciones
    var publicacionesContainer = document.getElementById('publicacionesContainer');

    // Limpiar el contenido anterior (si lo hay)
    publicacionesContainer.innerHTML = '';

    // Iterar sobre cada publicación y crear elementos HTML dinámicamente
    publicaciones.forEach(function(publicacion, index) {
        // Crear un contenedor para cada publicación
        var publicacionDiv = document.createElement('div');
        publicacionDiv.className = 'publicacion';

        // Agregar título al contenedor
        var titulo = document.createElement('h2');
        titulo.textContent = publicacion.title;
        publicacionDiv.appendChild(titulo);

        // Agregar descripción al contenedor
        var descripcion = document.createElement('p');
        descripcion.textContent = publicacion.description;
        publicacionDiv.appendChild(descripcion);

        // Agregar botón de administrar publicación
        var botonAdministrar = document.createElement('button');
        botonAdministrar.textContent = 'Administrar Publicación';
        botonAdministrar.addEventListener('click', function() {
            verDetalle(publicacion);
        });
        publicacionDiv.appendChild(botonAdministrar);

        // Agregar el contenedor al elemento principal
        publicacionesContainer.appendChild(publicacionDiv);

        // Agregar una línea divisoria después de cada elemento, excepto el último
        if (index < publicaciones.length - 1) {
            var lineaDivisoria = document.createElement('hr');
            publicacionesContainer.appendChild(lineaDivisoria);
        }
    });
}


function verDetalle(publicacion) {
    // Guardar el resultado actual en el almacenamiento local para que pueda ser recuperado en publicacion.html
    localStorage.setItem('publicacion', JSON.stringify(publicacion));

    // Redireccionar a la página publicacion.html
    window.location.href = 'publicacion.html';
}

function volver(){
    window.location.href = 'vender.html';
}
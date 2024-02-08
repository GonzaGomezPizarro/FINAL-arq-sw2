// publicacion.js

// Recuperar el resultado almacenado en el almacenamiento local
var publicacion = JSON.parse(localStorage.getItem('publicacion'));

console.log(publicacion)

// Obtener el elemento principal donde se mostrarán los detalles
var detalleContainer = document.getElementById('detalleContainer');

// Mostrar las fotos de la propiedad
var fotosContainer = document.createElement('div');
fotosContainer.className = 'fotos-container';

publicacion.photos.forEach(function(foto) {
    var imagen = document.createElement('img');
    imagen.src = foto;
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

function eliminarPublicacion(){

}
// Función para mostrar un formulario de edición
function editar() {
    // Obtener el elemento principal donde se mostrará el formulario
    var detalleContainer = document.getElementById('detalleContainer');

    // Limpiar el contenido anterior (si lo hay)
    detalleContainer.innerHTML = '';

    // Crear el formulario
    var formulario = document.createElement('form');

    // Iterar sobre cada campo y crear elementos HTML dinámicamente
    campos.forEach(function(campo) {
        // Crear un contenedor para cada campo
        var campoContainer = document.createElement('div');
        campoContainer.className = 'campo-edicion';

        // Crear etiqueta para el campo
        var etiqueta = document.createElement('label');
        etiqueta.textContent = `${campo.charAt(0).toUpperCase() + campo.slice(1)}: `;
        campoContainer.appendChild(etiqueta);

        // Crear campo de entrada para el formulario
        var inputCampo = document.createElement('input');
        inputCampo.type = 'text';
        inputCampo.id = `${campo}Input`;
        inputCampo.value = publicacion[campo];
        campoContainer.appendChild(inputCampo);

        // Agregar el contenedor al formulario
        formulario.appendChild(campoContainer);
    });

    // Botón de guardar cambios
    var botonGuardar = document.createElement('button');
    botonGuardar.textContent = 'Guardar Cambios';
    botonGuardar.type = 'button'; // Evitar que el formulario se envíe al hacer clic en el botón
    botonGuardar.addEventListener('click', function() {
        guardarCambios();
    });
    formulario.appendChild(botonGuardar);

    // Agregar el formulario al elemento principal
    detalleContainer.appendChild(formulario);
}
// Función para guardar cambios
function guardarCambios() {
    // Obtener el elemento principal donde se mostrará el formulario
    var detalleContainer = document.getElementById('detalleContainer');

    // Recopilar los nuevos valores del formulario
    var nuevosValores = {};
    campos.forEach(function(campo) {
        var inputCampo = document.getElementById(`${campo}Input`);
        nuevosValores[campo] = inputCampo.value;
    });

    // Comparar si los nuevos valores son diferentes de los valores originales
    if (JSON.stringify(nuevosValores) === JSON.stringify(publicacion)) {
        // No hay cambios, regresar
        return;
    }

    // Realizar una solicitud HTTP DELETE para eliminar el ítem existente
    fetch(`http://localhost:8091/item/${publicacion.id}`, {
        method: 'DELETE',
    })
    .then(response => response.json())
    .then(data => {
        console.log('Ítem eliminado con éxito:', data);

        // Realizar una solicitud HTTP POST para crear uno nuevo con los nuevos valores
        fetch('http://localhost:8091/item', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(nuevosValores)
        })
        .then(response => response.json())
        .then(data => {
            console.log('Ítem creado con nuevos valores:', data);

            // Puedes agregar aquí lógica adicional si es necesario

            // Redirigir a la página de publicaciones después de guardar cambios
            window.location.href = 'publicaciones.html';
        })
        .catch(error => {
            console.error('Error al crear el nuevo ítem:', error);
        });
    })
    .catch(error => {
        console.error('Error al eliminar el ítem existente:', error);
    });
}

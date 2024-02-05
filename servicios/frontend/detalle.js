// detalle.js

// Recuperar el resultado almacenado en el almacenamiento local
var detalleResultado = JSON.parse(localStorage.getItem('detalleResultado'));

// Obtener el elemento principal donde se mostrarán los detalles
var detalleContainer = document.getElementById('detalleContainer');

// Mostrar las fotos de la propiedad
var fotosContainer = document.createElement('div');
fotosContainer.className = 'fotos-container';

detalleResultado.photos.forEach(function(foto) {
    var imagen = document.createElement('img');
    imagen.src = foto;
    fotosContainer.appendChild(imagen);
});

detalleContainer.appendChild(fotosContainer);

// Mostrar los detalles de la propiedad
var campos = ['title', 'description', 'country', 'state', 'city', 'address', 'price', 'bedrooms', 'bathrooms', 'mts2', 'userId'];

campos.forEach(function(campo) {
    var etiqueta = document.createElement('p');
    etiqueta.textContent = `${campo.charAt(0).toUpperCase() + campo.slice(1)}: ${detalleResultado[campo]}`;
    detalleContainer.appendChild(etiqueta);
});

// Función para enviar mensajes
function enviarMensaje() {
    var mensajeInput = document.getElementById('mensajeInput');
    var mensaje = mensajeInput.value;

    // Verificar que haya un mensaje antes de enviar la solicitud
    if (mensaje) {
        // Construir el objeto JSON con la información requerida
        var mensajeData = {
            "content": mensaje,
            "receiver": detalleResultado.userId,
            "item": detalleResultado.id
        };

        // Realizar la solicitud HTTP POST
        fetch('http://localhost:8092/message', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(mensajeData)
        })
        .then(response => response.json())
        .then(data => {
            console.log('Mensaje enviado con éxito:', data);
            // Puedes agregar aquí lógica adicional si es necesario
        })
        .catch(error => {
            console.error('Error al enviar el mensaje:', error);
        });

        // Limpia el campo de entrada después de enviar el mensaje
        mensajeInput.value = '';
    }
}

// Limpiar el resultado almacenado en el almacenamiento local después de mostrar los detalles
localStorage.removeItem('detalleResultado');

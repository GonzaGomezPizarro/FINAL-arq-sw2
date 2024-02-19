// detalle.js

// Recuperar el resultado almacenado en el almacenamiento local
var detalleResultado = JSON.parse(localStorage.getItem('detalleResultado'));

console.log(detalleResultado)
console.log(detalleResultado.id)

// Obtener el elemento principal donde se mostrarán los detalles
var detalleContainer = document.getElementById('detalleContainer');

// Mostrar las fotos de la propiedad
var fotosContainer = document.createElement('div');
fotosContainer.className = 'imagenes';

detalleResultado.photos.forEach(function(fotoBase64) {
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
            "content": String(mensaje),
            "receiver": detalleResultado.userId,
            "item": String(detalleResultado.id) 
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
    setTimeout(function() {
        // Llama a la función getMessages después de 2 segundos
        getMessages(detalleResultado.id);
    }, 2000);
}

// Limpiar el resultado almacenado en el almacenamiento local después de mostrar los detalles
//localStorage.removeItem('detalleResultado');

function getMessages(itemId) {
    // Construir la URL de la solicitud
    itemId = String(itemId)
    var url = `http://localhost:8092/messagesByItem/${itemId}`;

    // Realizar la solicitud GET utilizando Fetch
    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log("Mensajes relacionados con el ítem:", data);
            if (data == null) { return; }
            // Mostrar los mensajes en el contenedor de todos los mensajes
            mostrarTodosMensajes(data);
        })
        .catch(error => {
            console.error("Error al obtener mensajes:", error);
        });
}


// Función para mostrar todos los mensajes en el contenedor correspondiente
function mostrarTodosMensajes(mensajes) {
    // Obtener el contenedor de mensajes
    var mensajeContainer = document.getElementById('todosMensajesContainer');

    // Limpiar el contenido anterior (si lo hay)
    mensajeContainer.innerHTML = '';

    // Iterar sobre cada mensaje y crear elementos HTML dinámicamente
    mensajes.forEach(function(mensaje, index) {
        // Crear un contenedor para cada mensaje
        var mensajeDiv = document.createElement('div');
        mensajeDiv.className = 'mensaje';

        // Agregar contenido del mensaje al contenedor
        var contenidoMensaje = document.createElement('p');
        contenidoMensaje.textContent = mensaje.content;
        mensajeDiv.appendChild(contenidoMensaje);

        // Agregar el contenedor al elemento principal
        mensajeContainer.appendChild(mensajeDiv);

        // Agregar una línea divisoria después de cada elemento, excepto el último
        if (index < mensajes.length - 1) {
            var lineaDivisoria = document.createElement('hr');
            mensajeContainer.appendChild(lineaDivisoria);
        }
    });
}

getMessages(detalleResultado.id);

function volver(){
    window.location.href = 'index.html';
}
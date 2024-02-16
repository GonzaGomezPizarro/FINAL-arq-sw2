document.addEventListener('DOMContentLoaded', function() {
    const fileInput = document.getElementById('fileInput');
    const buttonEnviar = document.getElementById('enviarButton'); // Corregido: Seleccionamos el botón por su ID

    buttonEnviar.addEventListener('click', function() {
        const file = fileInput.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = function(event) {
                const json = event.target.result;
                enviarJSON(json);
            };
            reader.readAsText(file);
        } else {
            console.error('No se ha seleccionado ningún archivo.');
        }
    });
});

function enviarJSON(json) {
    fetch('http://localhost:8095/items', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: json
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ocurrió un error al enviar el archivo.');
        }
        return response.json();
    })
    .then(data => {
        console.log('JSON enviado con éxito:', data);
        mostrarMensaje('JSON enviado con éxito.');
    })
    .catch(error => {
        console.error('Error al enviar el archivo:', error);
        mostrarMensaje('Error al enviar el archivo. Por favor, inténtalo de nuevo.');
    });
}

function mostrarMensaje(mensaje) {
    const resultadoContainer = document.getElementById('resultadoContainer');
    resultadoContainer.textContent = mensaje;
}

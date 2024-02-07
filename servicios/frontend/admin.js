function cargarArchivo() {
    const fileInput = document.getElementById('fileInput');
    
    const archivo = fileInput.files[0];

    if (archivo) {
        const lector = new FileReader();

        lector.onload = function(e) {
            const contenido = e.target.result;
            const json = JSON.parse(contenido);

            // Realizar la solicitud POST a http://localhost:8091/items
            realizarSolicitudPost(json);
        };

        lector.readAsText(archivo);
    } else {
        console.log('No se seleccionó ningún archivo.');
    }
}


function realizarSolicitudPost(data) {
    fetch('http://localhost:8091/items', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(responseData => {
        // Mostrar la respuesta del servidor en la página
        console.log(responseData)
        mostrarRespuesta(responseData);
    })
    .catch(error => {
        console.error('Error al realizar la solicitud POST:', error);
        mostrarRespuesta({ success: false, message: 'Error al realizar la solicitud POST: ' + error.message });
    });
}

function mostrarRespuesta(responseData) {
    const resultadoContainer = document.getElementById('resultadoContainer');
    resultadoContainer.innerHTML = '';

    const resultadoDiv = document.createElement('div');
    resultadoDiv.className = responseData.success ? 'success' : 'error';
    resultadoDiv.textContent = responseData.message;


    resultadoContainer.appendChild(resultadoDiv);

    console.log(resultadoDiv);
}
var contenedores = []; // Usar [] en lugar de new Array() es una forma más común de inicializar un array en JavaScript

function obtenerContenedores() {
    fetch('http://localhost:9000/contenedores')
        .then(response => response.json())
        .then(data => {
            // Manejar la respuesta de los contenedores
            contenedores = data;
            console.log(contenedores);
            mostrarContenedores(contenedores);
        })
        .catch(error => console.error('Error al obtener contenedores:', error));
}

function mostrarContenedores(contenedores) {
    // Obtener el elemento donde se mostrarán los contenedores
    var contenedorInfo = document.getElementById("contenedor-info");
    
    // Limpiar cualquier contenido previo
    contenedorInfo.innerHTML = "";
    
    // Crear elementos para cada contenedor y agregarlos al contenedor-info
    contenedores.forEach(contenedor => {
        var contenedorElemento = document.createElement("div");
        var separador = document.createElement("hr"); // Línea divisoria

        // Crear el texto con la información del contenedor
        var textoContenedor = document.createElement("span");
        textoContenedor.textContent = "Nombre: " + contenedor.name + ", ID: " + contenedor.id + ", Imagen: " + contenedor.imagen.name + ", Estado: " + contenedor.status;
        
        // Aplicar estilos al texto según el estado del contenedor
        if (contenedor.status.includes("Up")) {
            textoContenedor.classList.add("running-text");
        } else {
            textoContenedor.classList.add("stopped-text");
        }

        contenedorElemento.appendChild(textoContenedor);

        // Agregar un botón de acuerdo al estado del contenedor (Start o Stop)
        var botonEstado = document.createElement("button");
        if (contenedor.status.includes("Up")) {
            botonEstado.textContent = "Stop";
            botonEstado.classList.add("running"); // Agregar clase "running" si el contenedor está corriendo
            botonEstado.addEventListener("click", function() {
                stopContenedor(contenedor.id);
            });
        } else {
            botonEstado.textContent = "Start";
            botonEstado.classList.add("stopped"); // Agregar clase "stopped" si el contenedor está parado
            botonEstado.addEventListener("click", function() {
                startContenedor(contenedor.id);
            });
        }
        contenedorElemento.appendChild(botonEstado);

        // Agregar el botón "Delete"
        var botonDelete = document.createElement("button");
        botonDelete.textContent = "Delete";
        botonDelete.addEventListener("click", function() {
            borrarContenedor(contenedor.id);
        });
        contenedorElemento.appendChild(botonDelete);

        // Agregar el contenedor y el separador al contenedor-info
        contenedorInfo.appendChild(contenedorElemento);
        contenedorInfo.appendChild(separador);
    });
}

function obtenerImagenes() {
    fetch('http://localhost:9000/imagenes')
        .then(response => response.json())
        .then(data => {
            // Manejar la respuesta de las imágenes
            console.log(data);
            mostrarImagenes(data);
        })
        .catch(error => console.error('Error al obtener imágenes:', error));
}

function mostrarImagenes(imagenes) {
    // Obtener el elemento donde se mostrarán las imágenes
    var imagenesInfo = document.getElementById("imagen-info");
    
    // Limpiar cualquier contenido previo
    imagenesInfo.innerHTML = "";
    
    // Crear elementos para cada imagen y agregarlos al imagenes-info
    imagenes.forEach(imagen => {
        var imagenElemento = document.createElement("div");
        var separador = document.createElement("hr"); // Línea divisoria

        // Crear el texto con el nombre de la imagen
        var textoImagen = document.createTextNode("Nombre: " + imagen.name);
        imagenElemento.appendChild(textoImagen);

        // Agregar el elemento de imagen al contenedor de imágenes
        imagenesInfo.appendChild(imagenElemento);
        imagenesInfo.appendChild(separador);
    });
}

function startContenedor(id) {
    fetch(`http://localhost:9000/Pcontenedor/${id}`, {
        method: 'PUT'
    })
    .then(response => {
        if (response.ok) {
            // Si la respuesta es 200 (OK), mostrar un mensaje de alerta y actualizar la página
            alert("El contenedor se ha iniciado correctamente.");
            obtenerContenedores(); // Actualizar la lista de contenedores
        } else {
            // Si la respuesta no es 200, mostrar un mensaje de alerta indicando que no se pudo iniciar el contenedor
            alert("No se pudo iniciar el contenedor.");
        }
    })
    .catch(error => console.error('Error al iniciar contenedor:', error));
}

function stopContenedor(id) {
    fetch(`http://localhost:9000/Scontenedor/${id}`, {
        method: 'PUT'
    })
    .then(response => {
        if (response.ok) {
            // Si la respuesta es 200 (OK), mostrar un mensaje de alerta y actualizar la página
            alert("El contenedor se ha detenido correctamente.");
            obtenerContenedores(); // Actualizar la lista de contenedores
        } else {
            // Si la respuesta no es 200, mostrar un mensaje de alerta indicando que no se pudo detener el contenedor
            alert("No se pudo detener el contenedor.");
        }
    })
    .catch(error => console.error('Error al detener contenedor:', error));
}

// Función para mostrar el formulario de nuevo contenedor
function clickNuevoContenedor() {
    var formularioNuevoContenedor = document.getElementById("formulario-nuevo-contenedor");
    formularioNuevoContenedor.style.display = "block";
}

function crearContenedor(event) {
    // Evitar que el formulario se envíe automáticamente
    event.preventDefault();

    // Obtener los valores del formulario
    var nombre = document.getElementById("nombre").value;
    var nombreImagen = document.getElementById("nombre_imagen").value;
    var puertoInterno = document.getElementById("puerto_interno").value;
    var puertoExterno = document.getElementById("puerto_externo").value;
    
    // Verificar que el nombre de la imagen no esté vacío
    if (!nombreImagen.trim()) {
        alert("El nombre de la imagen es obligatorio.");
        return; // Salir de la función si el nombre de la imagen está vacío
    }
    
    // Crear el objeto JSON con los valores
    var nuevoContenedor = {
        "name": nombre || "",
        "imagen": {
            "name": nombreImagen
        }
    };
    
    // Agregar los puertos al objeto JSON si están presentes
    if (puertoInterno) {
        nuevoContenedor.internal_port = parseInt(puertoInterno);
    }
    if (puertoExterno) {
        nuevoContenedor.external_port = parseInt(puertoExterno);
    }

    // Realizar una solicitud POST al servidor para crear el contenedor
    fetch('http://localhost:9000/contenedor', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(nuevoContenedor)
    })
    .then(response => {
        if (response.ok) {
            // Si la respuesta es 200 (OK), mostrar un mensaje de alerta con la respuesta del servidor
            response.json().then(data => {
                alert("El contenedor se ha creado correctamente:\n" + JSON.stringify(data));
                // Cerrar el formulario
                var formularioNuevoContenedor = document.getElementById("formulario-nuevo-contenedor");
                formularioNuevoContenedor.style.display = "none";
                // Actualizar la lista de contenedores
                obtenerContenedores();
            });
        } else {
            // Si la respuesta no es 200, mostrar un mensaje de alerta indicando el error
            alert("Error al crear el contenedor. Por favor, revise los datos e intente nuevamente.");
        }
    })
    .catch(error => console.error('Error al crear contenedor:', error));
}


function borrarContenedor(id) {
    fetch(`http://localhost:9000/contenedor/${id}`, {
        method: 'DELETE'
    })
    .then(response => {
        if (response.ok) {
            // Si la respuesta es 200 (OK), mostrar un mensaje de alerta indicando que se ha borrado el contenedor
            alert("El contenedor se ha borrado correctamente.");
            // Actualizar la lista de contenedores
            obtenerContenedores();
        } else {
            // Si la respuesta no es 200, mostrar un mensaje de alerta indicando que no se pudo borrar el contenedor
            alert("No se pudo borrar el contenedor.");
        }
    })
    .catch(error => console.error('Error al borrar contenedor:', error));
}

obtenerContenedores();
obtenerImagenes();

setInterval(function() {
    obtenerContenedores();
    obtenerImagenes();
}, 30000); // 30 segundos en milisegundos


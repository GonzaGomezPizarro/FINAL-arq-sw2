var textobusqueda; 

function buscar() {
    // Obtener el valor del input con id 'searchInput'
    var inputElement = document.getElementById('searchInput');
    var textobusqueda = inputElement.value;

    if (textobusqueda == null || textobusqueda == "") {
        buscarTodo();
    } else {
        var url = "http://localhost:8000/search/" + textobusqueda;

        // Realizar la solicitud GET utilizando Fetch
        fetch(url)
            .then(response => response.json())
            .then(data => {
                console.log("Resultados de búsqueda:", data);
                mostrarResultados(data);
            })
            .catch(error => {
                console.error("Error al realizar la solicitud:", error);
            });
    }
}

function buscarTodo() {
    var url = "http://localhost:8000/searchAll";

    // Realizar la solicitud GET utilizando Fetch
    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log("Todos los resultados de búsqueda:", data);
            mostrarResultados(data);
        })
        .catch(error => {
            console.error("Error al realizar la solicitud:", error);
        });
}

function mostrarResultados(resultados) {
    // Obtener el elemento principal donde se mostrarán los resultados
    var resultadosContainer = document.getElementById('resultadosContainer');

    // Limpiar el contenido anterior (si lo hay)
    resultadosContainer.innerHTML = '';

    // Iterar sobre cada resultado y crear elementos HTML dinámicamente
    resultados.forEach(function(resultado, index) {
        // Crear un contenedor para cada resultado
        var resultadoDiv = document.createElement('div');
        resultadoDiv.className = 'resultado';

        // Agregar imágenes al contenedor
        resultado.photos.forEach(function(photo) {
            var imagen = document.createElement('img');
            imagen.src = photo;
            resultadoDiv.appendChild(imagen);
        });

        // Agregar título al contenedor
        var titulo = document.createElement('h2');
        titulo.textContent = resultado.title;
        resultadoDiv.appendChild(titulo);

        var descripcion = document.createElement('p');
        descripcion.textContent = resultado.description;
        resultadoDiv.appendChild(descripcion);

        // Agregar botón de detalle
        var botonDetalle = document.createElement('button');
        botonDetalle.textContent = 'Ver Detalle';
        botonDetalle.addEventListener('click', function() {
            verDetalle(resultado);
        });
        resultadoDiv.appendChild(botonDetalle);

        // Agregar el contenedor al elemento principal
        resultadosContainer.appendChild(resultadoDiv);

        // Agregar una línea divisoria después de cada elemento, excepto el último
        if (index < resultados.length - 1) {
            var lineaDivisoria = document.createElement('hr');
            resultadosContainer.appendChild(lineaDivisoria);
        }
    });
}

function verDetalle(resultado) {
    // Guardar el resultado actual en el almacenamiento local para que pueda ser recuperado en detalle.html
    localStorage.setItem('detalleResultado', JSON.stringify(resultado));

    // Redireccionar a la página detalle.html
    window.location.href = 'detalle.html';
}

function irAVender() {
    // Verificar si hay un token almacenado
    const token = localStorage.getItem('token');

    if (token) {
        // Redirigir a la página "vender.html" si hay un token
        window.location.href = 'vender.html';
    } else {
        // Redirigir a la página de inicio de sesión si no hay un token
        window.location.href = 'login.html';
    }
}

var textobusqueda;
var resultadosPorPagina = 10;
var paginaActual = 1;

var totalResultados = 0;

// Cargar catálogo inicial
buscarTodo();

function buscar() {
    var inputElement = document.getElementById('searchInput');
    textobusqueda = inputElement.value;

    if (!textobusqueda) {
        buscarTodo();
    } else {
        var url = "http://localhost:8000/search/" + textobusqueda;
        cargarResultadosPorPagina(url, paginaActual);
    }
}

function buscarTodo() {
    paginaActual = 1; // Reiniciar la página actual
    var url = "http://localhost:8000/searchAll";
    cargarResultadosPorPagina(url, paginaActual);
}

function cargarResultadosPorPagina(url, pagina) {
    var inicio = (pagina - 1) * resultadosPorPagina;
    var fin = inicio + resultadosPorPagina;

    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log("Resultados de búsqueda:", data);

            // Obtener el número total de resultados disponibles
            totalResultados = data.length;

            // Obtener los resultados de la página actual
            var resultadosPagina = data.slice(inicio, fin);

            // Mapear cada resultado con la función que maneja la carga de imágenes
            const resultadosConPromesas = resultadosPagina.map(resultado => {
                return cargarImagenes(resultado);
            });

            // Esperar hasta que todas las promesas se resuelvan
            Promise.all(resultadosConPromesas)
                .then(resultadosConImagenes => {
                    // Mostrar resultados con imágenes
                    mostrarResultados(resultadosConImagenes);
                })
                .catch(error => {
                    console.error("Error al cargar imágenes:", error);
                });
        })
        .catch(error => {
            console.error("Error al realizar la solicitud:", error);
        });
}

function irAPaginaSiguiente() {
    var totalPaginas = Math.ceil(totalResultados / resultadosPorPagina);
    if (paginaActual < totalPaginas) {
        paginaActual++;
        cargarResultadosPorPagina(url, paginaActual);
    }
}


function cargarImagenes(resultado) {
    const imagenesPromesas = resultado.photos.map(photo => {
        return new Promise((resolve, reject) => {
            const imagen = new Image();
            imagen.src = photo;
            imagen.onload = () => resolve(imagen);
            imagen.onerror = error => reject(error);
        });
    });

    // Retornar una promesa que se resuelve con el resultado y las imágenes cargadas
    return Promise.all(imagenesPromesas)
        .then(imagenesCargadas => {
            resultado.imagenesCargadas = imagenesCargadas;
            return resultado;
        });
}

function mostrarResultados(resultados) {
    var resultadosContainer = document.getElementById('resultadosContainer');
    resultadosContainer.innerHTML = '';

    resultados.forEach(function(resultado, index) {
        var resultadoDiv = document.createElement('div');
        resultadoDiv.className = 'resultado';

        // Agregar imágenes al contenedor
        resultado.imagenesCargadas.forEach(function(imagen) {
            resultadoDiv.appendChild(imagen);
        });

        var titulo = document.createElement('h2');
        titulo.textContent = resultado.title;
        resultadoDiv.appendChild(titulo);

        var descripcion = document.createElement('p');
        descripcion.textContent = resultado.description;
        resultadoDiv.appendChild(descripcion);

        var botonDetalle = document.createElement('button');
        botonDetalle.textContent = 'Ver Detalle';
        botonDetalle.addEventListener('click', function() {
            verDetalle(resultado);
        });
        resultadoDiv.appendChild(botonDetalle);

        resultadosContainer.appendChild(resultadoDiv);

        if (index < resultados.length - 1) {
            var lineaDivisoria = document.createElement('hr');
            resultadosContainer.appendChild(lineaDivisoria);
        }
    });
}

function verDetalle(resultado) {
    localStorage.setItem('detalleResultado', JSON.stringify(resultado));
    window.location.href = 'detalle.html';
}

function irAPaginaAnterior() {
    if (paginaActual > 1) {
        paginaActual--;
        buscar();
    }
}

function irAVender() {
    const token = localStorage.getItem('token');

    if (token) {
        window.location.href = 'vender.html';
    } else {
        window.location.href = 'login.html';
    }
}

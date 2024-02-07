var textobusqueda;

// Cargar catálogo inicial
buscarTodo();

function buscar() {
    var inputElement = document.getElementById('searchInput');
    textobusqueda = inputElement.value;

    if (!textobusqueda) {
        buscarTodo();
    } else {
        var url = "http://localhost:8000/search/" + textobusqueda;
        cargarResultados(url);
    }
}

function buscarTodo() {
    var url = "http://localhost:8000/searchAll";
    cargarResultados(url);
}


function cargarResultados(url) {
    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log("Resultados de búsqueda:", data);

            // Mapear cada resultado con la función que maneja la carga de imágenes
            const resultadosConPromesas = data.map(resultado => {
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

// Función para cargar imágenes de un resultado
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

        // Verificar si todas las imágenes están cargadas antes de agregarlas al contenedor
        if (resultado.imagenesCargadas.every(imagen => imagen.complete)) {
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
        } else {
            // Si no todas las imágenes están cargadas, espera y vuelve a intentar
            setTimeout(function() {
                mostrarResultados(resultados);
            }, 100);
        }
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

function irAVender() {
    const token = localStorage.getItem('token');

    if (token) {
        window.location.href = 'vender.html';
    } else {
        window.location.href = 'login.html';
    }
}

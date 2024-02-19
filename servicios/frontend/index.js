var textobusqueda;
var totalResultados = 0;

console.log("Total resultados");
// Cargar catálogo inicial
buscarTodo();

function buscar() {
    var inputElement = document.getElementById('searchInput');
    textobusqueda = inputElement.value;

    if (!textobusqueda) {
        buscarTodo();
    } else {
        BuscarTextoBusqueda(textobusqueda);
    }
}

function buscarTodo() {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "http://localhost:8000/searchAll", true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var data = JSON.parse(xhr.responseText);
                console.log(data);
                alert("Búsqueda realizada correctamente.");
                mostrarResultados(data);
            } else {
                alert("Error al realizar la búsqueda.");
            }
        }
    };
    xhr.send();
}

function BuscarTextoBusqueda(textobusqueda) {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "http://localhost:8000/search/" + textobusqueda, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var data = JSON.parse(xhr.responseText);
                console.log(data);
                alert("Búsqueda realizada correctamente.");
                mostrarResultados(data);
                
            } else {
                alert("Error al realizar la búsqueda.");
            }
        }
    };
    xhr.send();
}

// Función para mostrar los resultados con imágenes cargadas de forma asíncrona
function mostrarResultados(resultados) {
    console.log("Mostrando resultados...");
    var resultadosContainer = document.getElementById('resultadosContainer');
    resultadosContainer.innerHTML = '';

    resultados.forEach(function(resultado, index) {
        var resultadoDiv = document.createElement('div');
        resultadoDiv.className = 'resultado';

        var titulo = document.createElement('h2');
        titulo.textContent = resultado.title;
        resultadoDiv.appendChild(titulo);

        var imagenesDiv = document.createElement('div');
        imagenesDiv.className = 'imagenes'; // Contenedor de imágenes
        resultadoDiv.appendChild(imagenesDiv);

        var descripcion = document.createElement('p');
        descripcion.textContent = resultado.description;
        resultadoDiv.appendChild(descripcion);

        var infoHabitaciones = document.createElement('p');
        infoHabitaciones.textContent = 'Bedrooms: ' + resultado.bedrooms + ' - Baños: ' + resultado.bathrooms;
        resultadoDiv.appendChild(infoHabitaciones);

        var verDetalleBtn = document.createElement('button');
        verDetalleBtn.textContent = 'Ver Detalle';
        verDetalleBtn.addEventListener('click', function() {
            verDetalle(resultado);
        });
        resultadoDiv.appendChild(verDetalleBtn);

        resultadosContainer.appendChild(resultadoDiv);

        // Llamar a la función para cargar las imágenes de forma asíncrona
        cargarImagenesAsincronas(resultado.photos, imagenesDiv);
        
        if (index < resultados.length - 1) {
            var lineaDivisoria = document.createElement('hr');
            resultadosContainer.appendChild(lineaDivisoria);
        }
    });
}

// Función para cargar las imágenes de forma asíncrona
function cargarImagenesAsincronas(photos, imagenesDiv) {
    if (photos && photos.length > 0) {
        photos.forEach(function(fotoBase64) {
            var img = document.createElement('img');
            img.alt = 'Foto del resultado';
            img.className = 'imagen';
            imagenesDiv.appendChild(img);

            // Asignar el src de forma asíncrona
            setTimeout(function() {
                img.src = 'data:image/jpeg;base64,' + fotoBase64;
            }, 0);
        });
    } else {
        var sinImagenes = document.createElement('p');
        sinImagenes.textContent = 'No hay fotos disponibles.';
        imagenesDiv.appendChild(sinImagenes);
    }
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

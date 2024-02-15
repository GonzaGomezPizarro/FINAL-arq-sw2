// admin.js

let currentPage = 1;
const itemsPerPage = 10; // Número de elementos por página

function cargarArchivo() {
    // Tu función para cargar el archivo y realizar la solicitud POST
    // ...

    // Ejemplo: realizando la solicitud POST con datos ficticios
    const data = { message: 'Datos cargados exitosamente' };
    mostrarRespuesta(data);
}

function mostrarRespuesta(responseData) {
    const resultadoContainer = document.getElementById('resultadoContainer');
    resultadoContainer.innerHTML = `<div>${responseData.message}</div>`;

    // Lógica para la paginación
    const paginationContainer = document.getElementById('pagination');
    paginationContainer.innerHTML = '';
    const totalPages = Math.ceil(responseData.length / itemsPerPage);
    for (let i = 1; i <= totalPages; i++) {
        const button = document.createElement('button');
        button.textContent = i;
        button.addEventListener('click', () => {
            currentPage = i;
            mostrarPagina(responseData);
        });
        paginationContainer.appendChild(button);
    }
    // Mostrar la página actual
    mostrarPagina(responseData);
}

function mostrarPagina(responseData) {
    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const itemsToShow = responseData.slice(startIndex, endIndex);

    const resultadoContainer = document.getElementById('resultadoContainer');
    resultadoContainer.innerHTML = '';
    itemsToShow.forEach(item => {
        const div = document.createElement('div');
        div.textContent = item;
        resultadoContainer.appendChild(div);
    });

    // Marcar el botón de paginación activo
    const paginationButtons = document.querySelectorAll('#pagination button');
    paginationButtons.forEach((button, index) => {
        if (index + 1 === currentPage) {
            button.classList.add('active');
        } else {
            button.classList.remove('active');
        }
    });
}

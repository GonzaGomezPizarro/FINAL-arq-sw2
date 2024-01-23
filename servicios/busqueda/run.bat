@echo off

echo.
echo -------------------------------
echo -   Creando Contenedor SOLR   -
echo -------------------------------
echo.
docker-compose up -d

echo.
echo ----------------------------------------
echo -     ejecutando Servicio BUSQUEDA     -
echo ----------------------------------------
echo.
go run main.go

echo  Presione una tecla para finalizar...
pause>nul

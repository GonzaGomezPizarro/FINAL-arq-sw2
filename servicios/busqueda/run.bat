@echo off

echo.
echo ----------------------------------------
echo -   Creando Contenedor ELASTICSEARCH   -
echo ----------------------------------------
echo.
docker-compose up -d
echo.
docker ps

echo.
echo ----------------------------------------
echo -     ejecutando Servicio BUSQUEDA     -
echo ----------------------------------------
echo.
go run main.go

echo  Presione una tecla para finalizar...
pause>nul

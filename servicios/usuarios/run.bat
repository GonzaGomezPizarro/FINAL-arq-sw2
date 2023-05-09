@echo off

echo.
echo ------------------------------------
echo -     Creando Contenedor MySQL     -
echo ------------------------------------
echo.
docker-compose up -d

echo.
echo ----------------------------------------
echo -     ejecutando Servicio USUARIOS     -
echo ----------------------------------------
echo.
go run main.go

echo  Presione una tecla para finalizar...
pause>nul

@echo off

docker ps

echo.
echo ----------------------------------------
echo -     ejecutando Servicio BUSQUEDA     -
echo ----------------------------------------
echo.
go run main.go

echo  Presione una tecla para finalizar...
pause>nul

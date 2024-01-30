@echo off

echo.
echo ----------------------------------------
echo -     Corriendo contenedor RABBITMQ    -
echo ----------------------------------------
echo.
docker-compose up -d

echo  Presione una tecla para finalizar...
pause>nul

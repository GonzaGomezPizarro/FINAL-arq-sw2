@echo off

echo.
echo ---------------------------------------
echo -     Creando Contenedor RABBITMQ     -
echo ---------------------------------------
echo.
docker-compose up -d
echo.
docker ps
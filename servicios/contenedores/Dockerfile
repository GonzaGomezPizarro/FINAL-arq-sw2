# Etapa de compilación
FROM golang:1.18 AS builder

WORKDIR /app
COPY . .

# Compila la aplicación
RUN go build -o ejecutable .

# Etapa final
FROM alpine:latest

# Copia solo el ejecutable compilado desde la etapa anterior
COPY --from=builder /app/ejecutable /app/ejecutable

# Define el punto de entrada
ENTRYPOINT ["/app/ejecutable"]

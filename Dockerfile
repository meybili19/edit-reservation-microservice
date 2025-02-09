# Usa la imagen oficial de Golang
FROM golang:1.23.4

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Descargar dependencias
RUN go mod tidy

# Compilar el código
RUN go build -o edit-reservation cmd/main.go

# Exponer el puerto 8081
EXPOSE 4001

# Comando para ejecutar la aplicación
CMD ["./edit-reservation"]

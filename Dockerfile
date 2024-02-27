# Usa la imagen base Alpine con Go instalado
FROM golang:alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios para compilar y ejecutar el microservicio
COPY . .

# Compila la aplicación Go
RUN go build -o main .

# Expone el puerto en el que la aplicación estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación cuando el contenedor se inicia
CMD ["./main"]

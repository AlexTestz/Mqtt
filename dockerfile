# Usa una imagen base de Go
FROM golang:1.20-alpine

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia el código fuente de Go al contenedor
COPY . .

# Instala dependencias
RUN go mod tidy

# Compila la aplicación Go
RUN go build -o mqtt-app .

# Expone el puerto en caso de que se necesite (no obligatorio para este caso)
EXPOSE 1883

# Ejecuta la aplicación
CMD ["./mqtt-app"]

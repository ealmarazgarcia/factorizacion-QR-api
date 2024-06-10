# Usar la imagen base de GO
FROM golang:1.22-alpine

# Establecer el directorio para copiar
WORKDIR /app

# Copiar el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descargar módulos Go
RUN go mod download

# Copiar el código fuente
COPY . .

# Establece el directorio de trabajo donde se encuentra main.go
WORKDIR /app/cmd

# Compilar
RUN go build -o main .

# Exponer el puerto
EXPOSE 8080

# Correr la aplicación
CMD ["./main"]
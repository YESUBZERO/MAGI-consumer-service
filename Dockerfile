# Imagen base
FROM golang:1.23.5-alpine3.21

# Directorio de trabajo
WORKDIR /app

# Copiar dependencias y codigo fuente
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Compilar el binario
RUN go build -o consumer-service cmd/main.go

# Ejecutar el binario
CMD ["./consumer-service"]
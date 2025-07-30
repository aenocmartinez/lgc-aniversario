# Etapa 1: Compilación
FROM golang:1.23 as builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags netgo -ldflags="-s -w -extldflags '-static'" -o main .

# Etapa 2: Imagen final compatible con glibc
FROM debian:bullseye-slim

WORKDIR /app

# ✅ Instalar certificados TLS
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Copiar el binario compilado
COPY --from=builder /app/main .

# ✅ Copiar la carpeta html
COPY --from=builder /app/html /app/html

COPY ./resources /app/resources

RUN chmod +x /app/main


# Comando de ejecución
CMD ["/app/main"]

#!/bin/bash

set -e

CONTAINER_NAME="lgc-aniversario-container"

compile() {
    echo "🔄 Compilando código Go para Linux..."
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags netgo -ldflags="-s -w -extldflags '-static'" -o main .
}

# 💥 Forzar eliminación del contenedor y su log
remove_container_and_logs() {
    if docker inspect $CONTAINER_NAME >/dev/null 2>&1; then
        echo "🧹 Eliminando contenedor para limpiar logs..."
        docker stop $CONTAINER_NAME >/dev/null 2>&1 || true
        docker rm $CONTAINER_NAME >/dev/null 2>&1 || true
    fi
}

# 🆙 Levanta el contenedor sin reconstruir imagen
start_container() {
    echo "🚀 Levantando contenedor..."
    docker compose up -d
}

# 🚀 Flujo principal
case "$1" in
    --compile)
        compile
        remove_container_and_logs
        start_container
        ;;
    --build)
        echo "🔨 Reconstruyendo imagen..."
        compile
        docker compose down
        docker compose up -d --build
        ;;
    *)
        echo "♻️ Reiniciando contenedor..."
        docker stop $CONTAINER_NAME >/dev/null 2>&1 || true
        docker start $CONTAINER_NAME || docker compose up -d
        ;;
esac

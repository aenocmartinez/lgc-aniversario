#!/bin/bash

set -e  # Detiene la ejecución si hay errores

CONTAINER_NAME="lgcebd-container"

# Función para compilar correctamente en la máquina host
compile() {
    echo "🔄 Compilando código Go para Linux..."
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags="-s -w -extldflags '-static'" -o main .
}

# 🔄 Función para limpiar logs del contenedor
clear_logs() {
    echo "🧹 Limpiando logs del contenedor..."
    docker stop $CONTAINER_NAME >/dev/null 2>&1 || true
    docker rm $CONTAINER_NAME >/dev/null 2>&1 || true
    docker compose up -d
}

# Función para reiniciar el contenedor sin reconstruir la imagen
restart_container() {
    echo "♻️ Reiniciando contenedor..."
    docker stop $CONTAINER_NAME >/dev/null 2>&1 || true

    # 🔄 Limpiar logs del contenedor correctamente
    echo "🧹 Reiniciando logs del contenedor..."
    docker logs --tail 0 $CONTAINER_NAME >/dev/null 2>&1 || true
    truncate -s 0 $(docker inspect --format='{{.LogPath}}' $CONTAINER_NAME) 2>/dev/null || true

    docker start $CONTAINER_NAME || docker compose up -d
    echo "✅ Contenedor en ejecución."
}

# Verifica si se debe compilar o construir
case "$1" in
    --compile)
        compile
        clear_logs
        ;;
    --build)
        echo "🔨 Construyendo imagen y reiniciando..."
        compile  # Asegurar que el binario sea correcto antes de reconstruir
        docker compose down
        docker compose up -d --build
        ;;
    *)
        restart_container
        ;;
esac

#!/bin/bash

set -e  # Detener ejecución si hay un error

CONTAINER_NAME="lgc-aniversario-container"

# Verificar si Docker está instalado
if ! command -v docker &> /dev/null; then
    echo "Error: Docker no está instalado."
    exit 1
fi

# Verificar si el contenedor está en ejecución antes de detenerlo
if docker ps --format '{{.Names}}' | grep -q "^$CONTAINER_NAME$"; then
    echo "🛑 Deteniendo contenedor: $CONTAINER_NAME..."
    docker stop "$CONTAINER_NAME"
    echo "✅ Contenedor detenido correctamente."
else
    echo "⚠️  El contenedor $CONTAINER_NAME no está en ejecución."
fi

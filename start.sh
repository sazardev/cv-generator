#!/bin/bash

# Start script para Render
echo "🚀 Starting CV Generator application..."

# Verificar que el binario existe
if [ ! -f "./bin/main" ]; then
    echo "❌ Binary not found. Building..."
    go build -o bin/main ./cmd/server
fi

# Ejecutar la aplicación
echo "✅ Starting server on port ${PORT:-3000}..."
exec ./bin/main

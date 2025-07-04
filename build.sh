#!/bin/bash

# Build script para Render
echo "🔧 Installing dependencies..."
go mod download

echo "🏗️ Building application..."
go build -o bin/main ./cmd/server

echo "✅ Build completed successfully!"

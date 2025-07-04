# Usar imagen oficial de Go como base
FROM golang:1.21-alpine AS builder

# Instalar git para las dependencias
RUN apk add --no-cache git

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Etapa final - usar imagen alpine minimalista
FROM alpine:latest

# Instalar certificados SSL para HTTPS requests
RUN apk --no-cache add ca-certificates

# Crear directorio de la aplicación
WORKDIR /root/

# Copiar el binario compilado desde la etapa anterior
COPY --from=builder /app/main .

# Copiar archivos estáticos y plantillas
COPY --from=builder /app/web ./web

# Exponer puerto
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./main"]

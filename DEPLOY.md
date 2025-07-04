# 🚀 Guía de Despliegue en Render

Esta guía te ayudará a desplegar tu aplicación CV Generator en Render de forma gratuita.

## Prerequisitos

1. ✅ Cuenta en [GitHub](https://github.com)
2. ✅ Cuenta en [Render](https://render.com)
3. ✅ Repositorio Git con tu código

## Pasos para Desplegar

### 1. Preparar el Repositorio

Si aún no has subido tu código a GitHub:

```bash
# Inicializar git (si no lo has hecho)
git init

# Agregar todos los archivos
git add .

# Hacer commit
git commit -m "Initial commit with CV Generator app"

# Agregar el repositorio remoto (reemplaza con tu URL)
git remote add origin https://github.com/tu-usuario/cv-generator.git

# Subir el código
git push -u origin main
```

### 2. Conectar con Render

1. Ve a [Render Dashboard](https://dashboard.render.com)
2. Haz clic en **"New +"** → **"Web Service"**
3. Conecta tu cuenta de GitHub si no lo has hecho
4. Selecciona tu repositorio `cv-generator`

### 3. Configurar el Servicio

En la página de configuración de Render:

#### Configuración Básica
- **Name**: `cv-generator` (o el nombre que prefieras)
- **Environment**: `Go`
- **Region**: `Oregon (US West)` (recomendado para latencia)
- **Branch**: `main` o `master`

#### Configuración de Build y Start
- **Build Command**: `go build -o bin/main ./cmd/server`
- **Start Command**: `./bin/main`

#### Variables de Entorno
Render detectará automáticamente que es una app Go, pero puedes agregar:
- `PORT`: `10000` (Render usa este puerto por defecto)
- `RENDER`: `true` (para activar modo producción)

### 4. Desplegar

1. Haz clic en **"Create Web Service"**
2. Render automáticamente:
   - Clonará tu repositorio
   - Ejecutará `go mod download`
   - Compilará tu aplicación
   - La iniciará en el puerto especificado

### 5. Verificar el Despliegue

Una vez completado el despliegue:

1. Render te dará una URL como: `https://cv-generator-xxxx.onrender.com`
2. Visita la URL para verificar que funciona
3. Prueba generar un CV para asegurar que el PDF funciona correctamente

## Archivos de Configuración Incluidos

El proyecto ya incluye todos los archivos necesarios:

- ✅ `Procfile` - Define el comando de inicio
- ✅ `render.yaml` - Configuración específica de Render
- ✅ `build.sh` - Script de build personalizado
- ✅ `start.sh` - Script de inicio
- ✅ `Dockerfile` - Para despliegues con Docker (opcional)
- ✅ `.gitignore` - Excluye archivos innecesarios

## Características en Producción

### ✅ Optimizaciones Incluidas
- **Sin dependencias externas**: Solo Go puro, no requiere wkhtmltopdf
- **Templates optimizados**: Reload deshabilitado en producción
- **UTF-8 completo**: Soporte para acentos y caracteres especiales
- **Logs detallados**: Para debugging en producción
- **Health check**: Endpoint `/health` para monitoreo

### ✅ Configuración de Producción
- **Puerto dinámico**: Lee `PORT` de variables de entorno
- **Modo producción**: Detecta `RENDER` env var
- **CORS habilitado**: Para integraciones futuras
- **Recovery middleware**: Manejo robusto de errores

## Troubleshooting

### Error de Build
Si falla el build, verifica:
- ✅ `go.mod` y `go.sum` están commitados
- ✅ La versión de Go es compatible (1.21+)
- ✅ Todos los archivos fuente están incluidos

### Error de Runtime
Si la app no inicia:
- ✅ Verifica que `PORT` esté configurado
- ✅ Revisa los logs en Render Dashboard
- ✅ Asegúrate de que las rutas estáticas existan

### PDF no se genera
Si hay problemas con PDFs:
- ✅ Verifica que `gofpdf` esté en `go.mod`
- ✅ Revisa que los templates HTML estén incluidos
- ✅ Confirma que la codificación UTF-8 funciona

## Actualizaciones

Para actualizar tu app en producción:

```bash
# Hacer cambios en tu código
git add .
git commit -m "Update feature X"
git push origin main
```

Render automáticamente detectará los cambios y redesplegará.

## Monitoreo

Una vez desplegado, puedes:

1. **Ver logs**: En Render Dashboard → Tu servicio → Logs
2. **Monitorear**: Endpoint `/health` para status
3. **Métricas**: CPU, memoria y requests en el dashboard
4. **Alertas**: Configurar notificaciones de downtime

## Costo

- ✅ **Plan Free**: Incluido
- ✅ **Limitaciones**: 512MB RAM, sleep después de 15 min de inactividad
- ✅ **Para producción seria**: Considera plan pago ($7/mes) para evitar sleep

¡Tu CV Generator estará disponible 24/7 en Internet! 🎉

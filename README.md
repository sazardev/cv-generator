# CV Generator

Un generador de CV minimalista y monocromático con diseño estilo Notion, desarrollado en Go con Fiber. Permite crear CVs profesionales con vista previa web y exportación a PDF con coincidencia visual exacta.

## Características

- 🎨 **Diseño Monocromático**: Solo blanco, negro y tonalidades de grises, inspirado en Notion
- 📱 **Interfaz Minimalista**: UI limpia y moderna con tipografía Notion-style
- 👁️ **Vista Previa en Tiempo Real**: Preview idéntico al PDF final
- 📄 **Exportación PDF Nativa**: Genera PDFs con Go puro, sin dependencias externas
- ⚡ **Completamente Portátil**: Solo necesitas Go - no requiere instalaciones adicionales
- 🌍 **Soporte UTF-8 Completo**: Maneja acentos, ñ y caracteres especiales correctamente
- �️ **Selector de Idioma**: Cambia entre inglés y español en tiempo real
- 🔄 **Traducción Automática**: Headers y niveles de habilidad se traducen según el idioma seleccionado
- �🔧 **Sin Dependencias Externas**: No requiere wkhtmltopdf ni otros ejecutables
- 📋 **Formularios Dinámicos**: Agrega/elimina experiencias, educación, habilidades dinámicamente
- 🏷️ **SEO Optimizado**: Meta tags, Open Graph, y structured data
- 🎯 **Favicon SVG**: Icono minimalista y vectorial

## Tecnologías

- **Backend**: Go + Fiber
- **Frontend**: HTML5, CSS3 (Grid/Flexbox), JavaScript
- **PDF**: gofpdf (librería nativa de Go)
- **Tipografía**: ui-sans-serif, -apple-system (estilo Notion)
- **Codificación**: UTF-8 completo con soporte para caracteres latinos

## 🚀 Inicio rápido

### Windows
```bash
# Ejecuta el archivo batch
run.bat
```

### Linux/Mac
```bash
# Instala dependencias
go mod tidy

# Ejecuta la aplicación
go run cmd/server/main.go
```

Abre tu navegador en `http://localhost:3000`

## 📋 Uso

1. **Selecciona tu idioma**: Usa el selector en la esquina superior derecha (🇪🇸/🇺🇸)
2. **Completa tu información personal**: Nombre, email, teléfono, enlaces profesionales
3. **Agrega tu experiencia**: Empresas, cargos, fechas y descripciones
4. **Incluye tu educación**: Títulos, instituciones y logros
5. **Lista tus habilidades**: Con niveles opcionales de competencia (Básico, Intermedio, Avanzado, Experto)
6. **Especifica idiomas**: Idiomas que hablas y tu nivel
7. **Vista previa**: Revisa cómo se ve tu CV
8. **Exporta**: Descarga tu PDF profesional con traducción automática según el idioma seleccionado

## 🎨 Diseño

El diseño se centra en:
- **Monocromatismo**: Elegante paleta en blanco y negro
- **Tipografía limpia**: Fuentes profesionales y legibles
- **Espaciado perfecto**: Diseño balanceado y organizado
- **Minimalismo funcional**: Simple pero sofisticado

## Instalación y Uso

### Requisitos previos
- Go 1.24.4 o superior

### Ejecutar localmente

1. Clona el repositorio:
```bash
git clone <tu-repositorio>
cd cv-generator
```

2. Instala las dependencias:
```bash
go mod tidy
```

3. Ejecuta la aplicación:
```bash
go run cmd/server/main.go
```

4. Abre tu navegador y ve a `http://localhost:3000`

### Variables de entorno

- `PORT`: Puerto en el que se ejecutará el servidor (por defecto: 3000)

## Estructura del proyecto

```
cv-generator/
├── cmd/
│   └── server/
│       └── main.go          # Punto de entrada de la aplicación
├── internal/
│   ├── config/
│   │   └── config.go        # Configuración de la aplicación
│   ├── handlers/
│   │   └── cv.go           # Manejadores HTTP
│   ├── models/
│   │   └── cv.go           # Modelos de datos
│   └── services/
│       └── pdf.go          # Servicio de generación de PDF
├── web/
│   ├── static/
│   │   ├── styles.css      # Estilos CSS
│   │   └── app.js          # JavaScript de la aplicación
│   └── templates/
│       └── index.html      # Plantilla HTML principal
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints

- `GET /` - Página principal del formulario
- `POST /generate` - Genera y descarga el PDF del CV
- `GET /health` - Health check del servidor

## Tecnologías utilizadas

- **Backend**: Go, Fiber framework
- **Frontend**: HTML5, CSS3, JavaScript vanilla
- **PDF Generation**: gofpdf (librería nativa de Go)
- **Codificación**: UTF-8 completo con corrección automática de caracteres mal codificados
- **Styling**: CSS personalizado inspirado en Notion

## 🌍 Internacionalización

El generador de CV soporta completamente inglés y español:

### Traducción Automática
- **Headers de secciones**: EXPERIENCE/EXPERIENCIA, EDUCATION/EDUCACIÓN, SKILLS/HABILIDADES, etc.
- **Niveles de habilidad**: Basic/Básico, Intermediate/Intermedio, Advanced/Avanzado, Expert/Experto
- **Palabras comunes**: Present/Presente, at/en

### Funcionalidades i18n
- ✅ Selector de idioma en tiempo real en la interfaz
- ✅ El idioma seleccionado se aplica automáticamente al PDF generado
- ✅ Soporte completo para caracteres UTF-8 (acentos, ñ, etc.)
- ✅ Traducción bidireccional: español ↔ inglés
- ✅ Fallback inteligente si no se encuentra traducción

## Contribuir

1. Fork el proyecto
2. Crea tu feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push al branch (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 🚀 Despliegue en Producción

### Render (Recomendado)

Render es una plataforma moderna que facilita el despliegue de aplicaciones Go:

#### Preparación

1. **Crea una cuenta en [Render](https://render.com)**

2. **Conecta tu repositorio de GitHub**

3. **Configura el servicio web:**
   - **Build Command**: `go build -o bin/main ./cmd/server`
   - **Start Command**: `./bin/main`
   - **Environment**: Go
   - **Plan**: Free (para empezar)

#### Variables de Entorno

No necesitas configurar variables especiales, la aplicación ya está optimizada para Render:

- `PORT`: Se configura automáticamente por Render
- Sin dependencias externas requeridas

#### Configuración Automática

La aplicación incluye:
- ✅ `Dockerfile` optimizado para producción
- ✅ Scripts de build y start configurados
- ✅ Health check endpoint en `/health`
- ✅ Configuración de puerto dinámico
- ✅ Archivos estáticos servidos correctamente

#### Pasos de Despliegue

1. **Sube tu código a GitHub**
2. **Crea un nuevo Web Service en Render**
3. **Conecta tu repositorio**
4. **Render detectará automáticamente que es una aplicación Go**
5. **¡Despliega!**

Tu aplicación estará disponible en: `https://tu-app-name.onrender.com`

### Otras Plataformas

#### Railway
```bash
# Install Railway CLI
npm install -g @railway/cli

# Login y deploy
railway login
railway init
railway up
```

#### Heroku
```bash
# Crear Procfile
echo "web: ./bin/main" > Procfile

# Deploy
heroku create tu-app-name
git push heroku main
```

#### DigitalOcean App Platform
1. Conectar repositorio
2. Configurar: `go build -o bin/main ./cmd/server`
3. Start: `./bin/main`
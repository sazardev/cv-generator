# CV Generator

CV Generator es una aplicación web simple que permite a los usuarios crear y descargar sus CVs en formato PDF. Está construida usando Go y el framework web Fiber.

La aplicación web proporciona una interfaz fácil de usar donde los usuarios pueden ingresar su información personal, educación, experiencia laboral, habilidades y otros detalles relevantes. Una vez que el usuario ha completado el formulario, pueden generar una versión PDF de su CV.

El estilo es muy simple y minimalista, enfocándose en el contenido más que en el diseño. El PDF generado es limpio y profesional, adecuado para aplicaciones de trabajo.

## Características
- Formulario amigable para ingresar detalles del CV
- Generación de PDF usando Go's `html/template` y `github.com/jung-kurt/gofpdf`
- Diseño minimalista inspirado en Notion
- Fácil de desplegar y ejecutar en Google Cloud Run o cualquier otra plataforma en la nube
- Autocontenido, no requiere dependencias externas
- Soporta múltiples idiomas (actualmente solo inglés y español)
- Vista previa en tiempo real del CV
- Interfaz responsive y moderna

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
- **PDF Generation**: gofpdf
- **Styling**: CSS personalizado inspirado en Notion

## Contribuir

1. Fork el proyecto
2. Crea tu feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push al branch (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request
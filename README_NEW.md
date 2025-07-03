# CV Generator

CV Generator es una aplicación web simple que permite a los usuarios crear y descargar sus CVs en formato PDF. Está construida usando Go y el framework web Fiber con una interfaz ultra-limpia inspirada en Notion.

La aplicación web proporciona una interfaz fácil de usar donde los usuarios pueden ingresar su información personal, educación, experiencia laboral, habilidades y otros detalles relevantes. Una vez que el usuario ha completado el formulario, pueden ver una vista previa EXACTA de cómo se verá el PDF y luego descargarlo.

El estilo es extremadamente minimalista y profesional, enfocándose en el contenido más que en el diseño. El PDF generado es limpio y profesional, adecuado para aplicaciones de trabajo.

## ✨ Características

- **Interfaz Ultra-Limpia**: Diseño minimalista inspirado en Notion con tipografía Inter
- **Vista Previa Exacta**: La vista previa es IDÉNTICA al PDF final, con dimensiones A4 precisas
- **Generación PDF Profesional**: PDFs limpios usando `github.com/jung-kurt/gofpdf`
- **Formulario Dinámico**: Agrega/elimina secciones de experiencia, educación, habilidades e idiomas
- **Responsive**: Funciona perfectamente en desktop, tablet y móvil
- **Validación en Tiempo Real**: Validación de formularios y manejo de errores
- **Exportación Instantánea**: Descarga directa del PDF con nombre personalizado
- **Sin Dependencias Externas**: Autocontenido y fácil de desplegar

## 🛠 Tecnologías utilizadas

- **Backend**: Go 1.24.4, Fiber framework
- **Frontend**: HTML5, CSS3, JavaScript vanilla
- **PDF Generation**: gofpdf
- **Styling**: CSS Grid, Flexbox, variables CSS personalizadas
- **Icons**: Font Awesome 6.0
- **Typography**: Inter font family

## 📋 Requisitos previos

- Go 1.24.4 o superior

## 🚀 Instalación y Uso

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

## 📁 Estructura del proyecto

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
│   │   ├── styles.css      # Estilos CSS (Notion-inspired)
│   │   └── app.js          # JavaScript de la aplicación
│   └── templates/
│       └── index.html      # Plantilla HTML principal
├── go.mod
├── go.sum
└── README.md
```

## 🎯 API Endpoints

- `GET /` - Página principal del formulario
- `POST /generate` - Genera y descarga el PDF del CV
- `GET /health` - Health check del servidor

## 🎨 Características de Diseño

### Vista Previa Exacta
- Dimensiones A4 precisas (794px x 1123px)
- Tipografía Arial idéntica al PDF
- Colores RGB exactos que coinciden con el PDF
- Márgenes y espaciado preciso
- Formato de texto idéntico

### Interfaz Limpia
- Paleta de colores inspirada en Notion
- Micro-interacciones suaves
- Estados hover y focus bien definidos
- Feedback visual inmediato
- Animaciones sutiles y profesionales

## 📱 Responsive Design

- **Desktop**: Experiencia completa con vista previa a tamaño real
- **Tablet**: Layout adaptado con navegación touch-friendly
- **Mobile**: Interfaz optimizada con formularios apilados

## ⚡ Rendimiento

- **Carga Rápida**: CSS y JS minimalistas
- **Sin Dependencias Pesadas**: Solo Font Awesome para iconos
- **Generación PDF Eficiente**: Procesamiento rápido en el servidor
- **Caching**: Headers apropiados para recursos estáticos

## 🔧 Personalización

El diseño utiliza variables CSS personalizadas que pueden ser fácilmente modificadas:

```css
:root {
    --bg-primary: #ffffff;
    --bg-secondary: #f7f6f3;
    --text-primary: #37352f;
    --accent-blue: #2eaadc;
    /* ... más variables */
}
```

## 🚀 Despliegue

Esta aplicación está lista para ser desplegada en:
- **Google Cloud Run**
- **AWS Lambda** (con adaptador)
- **Heroku**
- **DigitalOcean App Platform**
- **Cualquier VPS** con Docker

## 🤝 Contribuir

1. Fork el proyecto
2. Crea tu feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push al branch (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para detalles.

## 🙏 Acknowledgments

- Inspirado en el diseño limpio de Notion
- Tipografía Inter de Google Fonts
- Iconos de Font Awesome

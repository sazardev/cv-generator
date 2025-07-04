package services

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
	"strings"

	"cv-generator/internal/models"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PDFService struct {
	template *template.Template
}

func NewPDFService() *PDFService {
	log.Println("🔧 Initializing PDF service...")

	// Load CV template
	tmplPath := filepath.Join("web", "templates", "cv_template.html")
	log.Printf("📄 Loading template from: %s", tmplPath)

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("⚠️ Template file not found, using embedded template: %v", err)
		// Fallback to embedded template if file not found
		tmpl = template.Must(template.New("cv").Parse(getEmbeddedTemplate()))
	} else {
		log.Println("✅ Template loaded successfully from file")
	}

	log.Println("✅ PDF service initialized")
	return &PDFService{
		template: tmpl,
	}
}

func (s *PDFService) GenerateCV(cv models.CV) ([]byte, error) {
	log.Println("🎨 Generating HTML from template...")

	// Generate HTML from template
	var htmlBuffer bytes.Buffer
	err := s.template.Execute(&htmlBuffer, cv)
	if err != nil {
		log.Printf("❌ Template execution failed: %v", err)
		return nil, err
	}

	htmlContent := htmlBuffer.String()
	log.Printf("✅ HTML generated successfully, length: %d characters", len(htmlContent))

	// Show a preview of the HTML (safely)
	previewLength := 500
	if len(htmlContent) < previewLength {
		previewLength = len(htmlContent)
	}
	log.Printf("📋 HTML preview (first %d chars): %s...", previewLength, htmlContent[:previewLength])

	// Create PDF from HTML using wkhtmltopdf
	log.Println("🔧 Initializing wkhtmltopdf...")

	// Set the path to wkhtmltopdf executable
	wkhtmltopdf.SetPath("C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe")
	log.Println("✅ wkhtmltopdf path configured")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Printf("❌ Failed to create PDF generator: %v", err)
		return nil, err
	}
	log.Println("✅ PDF generator created")

	// Configure PDF options for professional output
	log.Println("⚙️ Configuring PDF options...")
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	// Set proper margins to ensure consistent spacing across pages
	pdfg.MarginTop.Set(25)    // 25mm top margin
	pdfg.MarginBottom.Set(25) // 25mm bottom margin
	pdfg.MarginLeft.Set(25)   // 25mm left margin
	pdfg.MarginRight.Set(25)  // 25mm right margin
	log.Println("✅ PDF options configured")

	// Add the HTML page
	log.Println("📄 Adding HTML page to PDF generator...")
	page := wkhtmltopdf.NewPageReader(strings.NewReader(htmlContent))
	page.DisableSmartShrinking.Set(true)
	page.EnableLocalFileAccess.Set(true)
	page.LoadErrorHandling.Set("ignore")
	// Additional page options for better multi-page rendering
	page.PrintMediaType.Set(true)
	page.NoBackground.Set(false)
	page.MinimumFontSize.Set(0)
	pdfg.AddPage(page)
	log.Println("✅ HTML page added")

	// Generate PDF
	log.Println("🎯 Creating PDF...")
	err = pdfg.Create()
	if err != nil {
		log.Printf("❌ PDF creation failed: %v", err)
		return nil, err
	}
	log.Println("✅ PDF created successfully")

	pdfBytes := pdfg.Bytes()
	log.Printf("✅ PDF generation completed, size: %d bytes", len(pdfBytes))

	return pdfBytes, nil
}

// Fallback embedded template
func getEmbeddedTemplate() string {
	return `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        @page { size: A4; margin: 0; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', system-ui, sans-serif;
            color: #37352f;
            line-height: 1.6;
            font-size: 11pt;
            background: white;
            margin: 0;
            padding: 0;
        }
        .name { 
            font-size: 24pt; 
            font-weight: 700; 
            margin-bottom: 10pt; 
            page-break-after: avoid;
        }
        .contact-info { 
            font-size: 9pt; 
            color: #6f6f6f; 
            margin-bottom: 20pt; 
        }
        .section { 
            margin-bottom: 20pt; 
            page-break-inside: avoid;
            orphans: 2;
            widows: 2;
        }
        .section-title { 
            font-size: 11pt; 
            font-weight: 700; 
            text-transform: uppercase; 
            margin-bottom: 10pt;
            border-bottom: 1pt solid #e3e2e0;
            padding-bottom: 4pt;
            page-break-after: avoid;
        }
        .item { 
            margin-bottom: 12pt; 
            page-break-inside: avoid;
            orphans: 2;
            widows: 2;
        }
        .item-title { 
            font-weight: 600; 
            margin-bottom: 3pt; 
            page-break-after: avoid;
        }
        .item-subtitle { 
            font-size: 9pt; 
            color: #6f6f6f; 
            font-style: italic; 
        }
    </style>
</head>
<body>
    <h1 class="name">{{.PersonalInfo.FullName}}</h1>
    <div class="contact-info">
        {{.PersonalInfo.Email}}{{if .PersonalInfo.Phone}} • {{.PersonalInfo.Phone}}{{end}}{{if .PersonalInfo.Location}} • {{.PersonalInfo.Location}}{{end}}
    </div>
    {{if .PersonalInfo.Summary}}
    <div class="section">
        <h2 class="section-title">Summary</h2>
        <p>{{.PersonalInfo.Summary}}</p>
    </div>
    {{end}}
</body>
</html>`
}

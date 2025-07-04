package services

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"cv-generator/internal/models"

	"github.com/jung-kurt/gofpdf"
)

type PDFService struct{}

func NewPDFService() *PDFService {
	log.Println("🔧 Initializing PDF service with gofpdf...")
	log.Println("✅ PDF service initialized - no external dependencies required!")
	return &PDFService{}
}

// cleanText ensures text is properly encoded for PDF and fixes common encoding issues
func (s *PDFService) cleanText(text string) string {
	// Convert common problematic characters from bad UTF-8 encoding
	replacements := map[string]string{
		"Ã¡":           "á",
		"Ã©":           "é",
		"Ã­":           "í",
		"Ã³":           "ó",
		"Ãº":           "ú",
		"Ã±":           "ñ",
		"BÃ¡sico":      "Básico",
		"bÃ¡sico":      "básico",
		"BÃ sico":      "Básico",
		"bÃ sico":      "básico",
		"Basico":       "Básico",
		"basico":       "básico",
		"intermedio":   "Intermedio",
		"avanzado":     "Avanzado",
		"experto":      "Experto",
		"EducaciÃ³n":   "Educación",
		"ExperiÃ©ncia": "Experiencia",
		"HabilidÃ¡des": "Habilidades",
	}

	result := text
	for bad, good := range replacements {
		result = strings.ReplaceAll(result, bad, good)
	}

	// Remove any remaining null bytes or invalid characters
	result = strings.ReplaceAll(result, "\x00", "")
	result = strings.TrimSpace(result)

	return result
}

func (s *PDFService) GenerateCV(cv models.CV) ([]byte, error) {
	log.Println("🎨 Generating PDF with gofpdf...")

	// Create new PDF document with better UTF-8 handling
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Configure translator to handle non-ASCII characters
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	// Set margins
	pdf.SetMargins(25, 25, 25)
	pdf.SetAutoPageBreak(true, 25)

	// Add first page
	pdf.AddPage()

	// Define colors (monochromatic theme)
	textColor := map[string]int{"r": 55, "g": 53, "b": 47}         // #37352f
	lightTextColor := map[string]int{"r": 111, "g": 111, "b": 111} // #6f6f6f
	separatorColor := map[string]int{"r": 227, "g": 226, "b": 224} // #e3e2e0

	// Header - Name
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 18)
	cleanName := tr(s.cleanText(cv.PersonalInfo.FullName))
	pdf.CellFormat(0, 12, cleanName, "", 1, "L", false, 0, "")
	pdf.Ln(3)

	// Contact Information
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(lightTextColor["r"], lightTextColor["g"], lightTextColor["b"])

	var contactParts []string
	if cv.PersonalInfo.Email != "" {
		contactParts = append(contactParts, tr(s.cleanText(cv.PersonalInfo.Email)))
	}
	if cv.PersonalInfo.Phone != "" {
		contactParts = append(contactParts, tr(s.cleanText(cv.PersonalInfo.Phone)))
	}
	if cv.PersonalInfo.Location != "" {
		contactParts = append(contactParts, tr(s.cleanText(cv.PersonalInfo.Location)))
	}
	if cv.PersonalInfo.LinkedIn != "" {
		contactParts = append(contactParts, "LinkedIn: "+tr(s.cleanText(cv.PersonalInfo.LinkedIn)))
	}
	if cv.PersonalInfo.GitHub != "" {
		contactParts = append(contactParts, "GitHub: "+tr(s.cleanText(cv.PersonalInfo.GitHub)))
	}
	if cv.PersonalInfo.Website != "" {
		contactParts = append(contactParts, tr(s.cleanText(cv.PersonalInfo.Website)))
	}

	if len(contactParts) > 0 {
		contactInfo := strings.Join(contactParts, " • ")
		// Split long contact info into multiple lines if needed
		lines := s.splitText(pdf, contactInfo, 160)
		for _, line := range lines {
			pdf.CellFormat(0, 5, line, "", 1, "L", false, 0, "")
		}
	}

	pdf.Ln(5)

	// Add separator line
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(8)

	// Summary Section
	if cv.PersonalInfo.Summary != "" {
		s.addSection(pdf, tr, "RESUMEN", s.cleanText(cv.PersonalInfo.Summary), textColor, lightTextColor, separatorColor)
	}

	// Experience Section
	if len(cv.Experience) > 0 {
		s.addExperienceSection(pdf, tr, cv.Experience, textColor, lightTextColor, separatorColor)
	}

	// Education Section
	if len(cv.Education) > 0 {
		s.addEducationSection(pdf, tr, cv.Education, textColor, lightTextColor, separatorColor)
	}

	// Skills Section
	if len(cv.Skills) > 0 {
		s.addSkillsSection(pdf, tr, cv.Skills, textColor, lightTextColor, separatorColor)
	}

	// Languages Section
	if len(cv.Languages) > 0 {
		s.addLanguagesSection(pdf, tr, cv.Languages, textColor, lightTextColor, separatorColor)
	}

	// Generate PDF bytes
	buffer := &bytes.Buffer{}
	err := pdf.Output(buffer)
	if err != nil {
		log.Printf("❌ PDF generation failed: %v", err)
		return nil, err
	}

	buf := buffer.Bytes()
	log.Printf("✅ PDF generated successfully, size: %d bytes", len(buf))

	return buf, nil
}

func (s *PDFService) addSection(pdf *gofpdf.Fpdf, tr func(string) string, title, content string, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, tr(s.cleanText(title)), "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	// Section content
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])

	lines := s.splitText(pdf, tr(content), 160)
	for _, line := range lines {
		pdf.CellFormat(0, 5, line, "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)
}

func (s *PDFService) addExperienceSection(pdf *gofpdf.Fpdf, tr func(string) string, experiences []models.Experience, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, tr("EXPERIENCIA"), "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	for i, exp := range experiences {
		// Item title
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
		title := fmt.Sprintf("%s en %s", tr(s.cleanText(exp.Position)), tr(s.cleanText(exp.Company)))
		pdf.CellFormat(0, 5, title, "", 1, "L", false, 0, "")

		// Item subtitle (dates)
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(lightTextColor["r"], lightTextColor["g"], lightTextColor["b"])
		dateRange := tr(s.cleanText(exp.StartDate))
		if exp.EndDate != "" {
			dateRange += " - " + tr(s.cleanText(exp.EndDate))
		} else {
			dateRange += " - " + tr("Presente")
		}
		pdf.CellFormat(0, 4, dateRange, "", 1, "L", false, 0, "")

		// Description
		if exp.Description != "" {
			pdf.SetFont("Arial", "", 10)
			pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
			lines := s.splitText(pdf, tr(s.cleanText(exp.Description)), 160)
			for _, line := range lines {
				pdf.CellFormat(0, 4, line, "", 1, "L", false, 0, "")
			}
		}

		if i < len(experiences)-1 {
			pdf.Ln(3)
		}
	}
	pdf.Ln(5)
}

func (s *PDFService) addEducationSection(pdf *gofpdf.Fpdf, tr func(string) string, education []models.Education, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, tr("EDUCACIÓN"), "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	for i, edu := range education {
		// Item title
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
		title := fmt.Sprintf("%s - %s", tr(s.cleanText(edu.Degree)), tr(s.cleanText(edu.Institution)))
		pdf.CellFormat(0, 5, title, "", 1, "L", false, 0, "")

		// Item subtitle (dates)
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(lightTextColor["r"], lightTextColor["g"], lightTextColor["b"])
		dateRange := tr(s.cleanText(edu.StartDate))
		if edu.EndDate != "" {
			dateRange += " - " + tr(s.cleanText(edu.EndDate))
		} else {
			dateRange += " - " + tr("Presente")
		}
		pdf.CellFormat(0, 4, dateRange, "", 1, "L", false, 0, "")

		// Description
		if edu.Description != "" {
			pdf.SetFont("Arial", "", 10)
			pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
			lines := s.splitText(pdf, tr(s.cleanText(edu.Description)), 160)
			for _, line := range lines {
				pdf.CellFormat(0, 4, line, "", 1, "L", false, 0, "")
			}
		}

		if i < len(education)-1 {
			pdf.Ln(3)
		}
	}
	pdf.Ln(5)
}

func (s *PDFService) addSkillsSection(pdf *gofpdf.Fpdf, tr func(string) string, skills []models.Skill, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, tr("HABILIDADES"), "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])

	// Display skills in columns
	x := 25.0
	y := pdf.GetY()
	colWidth := 80.0
	lineHeight := 5.0

	for i, skill := range skills {
		if i > 0 && i%2 == 0 {
			y += lineHeight
			x = 25.0
		} else if i > 0 {
			x = 105.0
		}

		pdf.SetXY(x, y)

		skillText := tr(s.cleanText(skill.Name))
		if skill.Level != "" {
			skillText += " (" + tr(s.cleanText(skill.Level)) + ")"
		}

		pdf.CellFormat(colWidth, lineHeight, skillText, "", 0, "L", false, 0, "")
	}

	pdf.SetXY(25, y+lineHeight+5)
}

func (s *PDFService) addLanguagesSection(pdf *gofpdf.Fpdf, tr func(string) string, languages []string, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, tr("IDIOMAS"), "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])

	var cleanLanguages []string
	for _, lang := range languages {
		cleanLanguages = append(cleanLanguages, tr(s.cleanText(lang)))
	}

	languageText := strings.Join(cleanLanguages, " • ")
	lines := s.splitText(pdf, languageText, 160)
	for _, line := range lines {
		pdf.CellFormat(0, 5, line, "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)
}

func (s *PDFService) splitText(pdf *gofpdf.Fpdf, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine []string

	for _, word := range words {
		testLine := append(currentLine, word)
		testText := strings.Join(testLine, " ")

		textWidth := pdf.GetStringWidth(testText)

		if textWidth <= maxWidth {
			currentLine = testLine
		} else {
			if len(currentLine) > 0 {
				lines = append(lines, strings.Join(currentLine, " "))
			}
			currentLine = []string{word}
		}
	}

	if len(currentLine) > 0 {
		lines = append(lines, strings.Join(currentLine, " "))
	}

	return lines
}

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

func (s *PDFService) GenerateCV(cv models.CV) ([]byte, error) {
	log.Println("🎨 Generating PDF with gofpdf...")

	// Create new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

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
	pdf.CellFormat(0, 12, cv.PersonalInfo.FullName, "", 1, "L", false, 0, "")
	pdf.Ln(3)

	// Contact Information
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(lightTextColor["r"], lightTextColor["g"], lightTextColor["b"])

	var contactParts []string
	if cv.PersonalInfo.Email != "" {
		contactParts = append(contactParts, cv.PersonalInfo.Email)
	}
	if cv.PersonalInfo.Phone != "" {
		contactParts = append(contactParts, cv.PersonalInfo.Phone)
	}
	if cv.PersonalInfo.Location != "" {
		contactParts = append(contactParts, cv.PersonalInfo.Location)
	}
	if cv.PersonalInfo.LinkedIn != "" {
		contactParts = append(contactParts, "LinkedIn: "+cv.PersonalInfo.LinkedIn)
	}
	if cv.PersonalInfo.GitHub != "" {
		contactParts = append(contactParts, "GitHub: "+cv.PersonalInfo.GitHub)
	}
	if cv.PersonalInfo.Website != "" {
		contactParts = append(contactParts, cv.PersonalInfo.Website)
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
		s.addSection(pdf, "SUMMARY", cv.PersonalInfo.Summary, textColor, lightTextColor, separatorColor)
	}

	// Experience Section
	if len(cv.Experience) > 0 {
		s.addExperienceSection(pdf, cv.Experience, textColor, lightTextColor, separatorColor)
	}

	// Education Section
	if len(cv.Education) > 0 {
		s.addEducationSection(pdf, cv.Education, textColor, lightTextColor, separatorColor)
	}

	// Skills Section
	if len(cv.Skills) > 0 {
		s.addSkillsSection(pdf, cv.Skills, textColor, lightTextColor, separatorColor)
	}

	// Languages Section
	if len(cv.Languages) > 0 {
		s.addLanguagesSection(pdf, cv.Languages, textColor, lightTextColor, separatorColor)
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

func (s *PDFService) addSection(pdf *gofpdf.Fpdf, title, content string, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, title, "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	// Section content
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])

	lines := s.splitText(pdf, content, 160)
	for _, line := range lines {
		pdf.CellFormat(0, 5, line, "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)
}

func (s *PDFService) addExperienceSection(pdf *gofpdf.Fpdf, experiences []models.Experience, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, "EXPERIENCE", "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	for i, exp := range experiences {
		// Item title
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
		title := fmt.Sprintf("%s at %s", exp.Position, exp.Company)
		pdf.CellFormat(0, 5, title, "", 1, "L", false, 0, "")

		// Item subtitle (dates)
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(lightTextColor["r"], lightTextColor["g"], lightTextColor["b"])
		dateRange := exp.StartDate
		if exp.EndDate != "" {
			dateRange += " - " + exp.EndDate
		} else {
			dateRange += " - Present"
		}
		pdf.CellFormat(0, 4, dateRange, "", 1, "L", false, 0, "")

		// Description
		if exp.Description != "" {
			pdf.SetFont("Arial", "", 10)
			pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
			lines := s.splitText(pdf, exp.Description, 160)
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

func (s *PDFService) addEducationSection(pdf *gofpdf.Fpdf, education []models.Education, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, "EDUCATION", "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	for i, edu := range education {
		// Item title
		pdf.SetFont("Arial", "B", 10)
		pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
		title := fmt.Sprintf("%s - %s", edu.Degree, edu.Institution)
		pdf.CellFormat(0, 5, title, "", 1, "L", false, 0, "")

		// Item subtitle (dates)
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(lightTextColor["r"], lightTextColor["g"], lightTextColor["b"])
		dateRange := edu.StartDate
		if edu.EndDate != "" {
			dateRange += " - " + edu.EndDate
		} else {
			dateRange += " - Present"
		}
		pdf.CellFormat(0, 4, dateRange, "", 1, "L", false, 0, "")

		// Description
		if edu.Description != "" {
			pdf.SetFont("Arial", "", 10)
			pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
			lines := s.splitText(pdf, edu.Description, 160)
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

func (s *PDFService) addSkillsSection(pdf *gofpdf.Fpdf, skills []models.Skill, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, "SKILLS", "", 1, "L", false, 0, "")

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

		skillText := skill.Name
		if skill.Level != "" {
			skillText += " (" + skill.Level + ")"
		}

		pdf.CellFormat(colWidth, lineHeight, skillText, "", 0, "L", false, 0, "")
	}

	pdf.SetXY(25, y+lineHeight+5)
}

func (s *PDFService) addLanguagesSection(pdf *gofpdf.Fpdf, languages []string, textColor, lightTextColor, separatorColor map[string]int) {
	// Section title
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, "LANGUAGES", "", 1, "L", false, 0, "")

	// Section separator
	pdf.SetDrawColor(separatorColor["r"], separatorColor["g"], separatorColor["b"])
	pdf.Line(25, pdf.GetY(), 185, pdf.GetY())
	pdf.Ln(3)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(textColor["r"], textColor["g"], textColor["b"])

	languageText := strings.Join(languages, " • ")
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

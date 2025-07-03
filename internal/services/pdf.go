package services

import (
	"bytes"
	"fmt"
	"strings"

	"cv-generator/internal/models"

	"github.com/jung-kurt/gofpdf"
)

type PDFService struct{}

func NewPDFService() *PDFService {
	return &PDFService{}
}

func (s *PDFService) GenerateCV(cv models.CV) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set margins
	pdf.SetMargins(20, 20, 20)
	pdf.SetAutoPageBreak(true, 20)

	// Header - Personal Information
	s.addHeader(pdf, cv.PersonalInfo)

	// Summary
	if cv.PersonalInfo.Summary != "" {
		s.addSection(pdf, "SUMMARY", cv.PersonalInfo.Summary)
	}

	// Experience
	if len(cv.Experience) > 0 {
		s.addExperienceSection(pdf, cv.Experience)
	}

	// Education
	if len(cv.Education) > 0 {
		s.addEducationSection(pdf, cv.Education)
	}

	// Skills
	if len(cv.Skills) > 0 {
		s.addSkillsSection(pdf, cv.Skills)
	}

	// Languages
	if len(cv.Languages) > 0 {
		s.addLanguagesSection(pdf, cv.Languages)
	}

	// Get PDF bytes
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	return buf.Bytes(), err
}

func (s *PDFService) addHeader(pdf *gofpdf.Fpdf, info models.PersonalInfo) {
	// Name
	pdf.SetFont("Arial", "B", 24)
	pdf.SetTextColor(44, 62, 80)
	pdf.Cell(0, 15, info.FullName)
	pdf.Ln(12)

	// Contact information
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(100, 100, 100)

	contacts := []string{}
	if info.Email != "" {
		contacts = append(contacts, info.Email)
	}
	if info.Phone != "" {
		contacts = append(contacts, info.Phone)
	}
	if info.Location != "" {
		contacts = append(contacts, info.Location)
	}
	if info.LinkedIn != "" {
		contacts = append(contacts, "LinkedIn: "+info.LinkedIn)
	}
	if info.GitHub != "" {
		contacts = append(contacts, "GitHub: "+info.GitHub)
	}
	if info.Website != "" {
		contacts = append(contacts, "Website: "+info.Website)
	}

	if len(contacts) > 0 {
		pdf.Cell(0, 5, strings.Join(contacts, " | "))
		pdf.Ln(8)
	}

	// Add separator line
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(20, pdf.GetY(), 190, pdf.GetY())
	pdf.Ln(8)
}

func (s *PDFService) addSection(pdf *gofpdf.Fpdf, title, content string) {
	// Section title
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(44, 62, 80)
	pdf.Cell(0, 8, title)
	pdf.Ln(6)

	// Content
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(80, 80, 80)

	// Split content into lines to handle long text
	lines := s.splitText(pdf, content, 170)
	for _, line := range lines {
		pdf.Cell(0, 5, line)
		pdf.Ln(5)
	}
	pdf.Ln(4)
}

func (s *PDFService) addExperienceSection(pdf *gofpdf.Fpdf, experiences []models.Experience) {
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(44, 62, 80)
	pdf.Cell(0, 8, "EXPERIENCE")
	pdf.Ln(6)

	for _, exp := range experiences {
		// Position and Company
		pdf.SetFont("Arial", "B", 12)
		pdf.SetTextColor(60, 60, 60)
		pdf.Cell(0, 6, fmt.Sprintf("%s at %s", exp.Position, exp.Company))
		pdf.Ln(5)

		// Dates
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(120, 120, 120)
		dateRange := fmt.Sprintf("%s - %s", exp.StartDate, exp.EndDate)
		if exp.EndDate == "" {
			dateRange = fmt.Sprintf("%s - Present", exp.StartDate)
		}
		pdf.Cell(0, 4, dateRange)
		pdf.Ln(4)

		// Description
		if exp.Description != "" {
			pdf.SetFont("Arial", "", 10)
			pdf.SetTextColor(80, 80, 80)
			lines := s.splitText(pdf, exp.Description, 170)
			for _, line := range lines {
				pdf.Cell(0, 4, line)
				pdf.Ln(4)
			}
		}
		pdf.Ln(3)
	}
	pdf.Ln(2)
}

func (s *PDFService) addEducationSection(pdf *gofpdf.Fpdf, educations []models.Education) {
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(44, 62, 80)
	pdf.Cell(0, 8, "EDUCATION")
	pdf.Ln(6)

	for _, edu := range educations {
		// Degree and Institution
		pdf.SetFont("Arial", "B", 12)
		pdf.SetTextColor(60, 60, 60)
		pdf.Cell(0, 6, fmt.Sprintf("%s - %s", edu.Degree, edu.Institution))
		pdf.Ln(5)

		// Dates
		pdf.SetFont("Arial", "I", 9)
		pdf.SetTextColor(120, 120, 120)
		dateRange := fmt.Sprintf("%s - %s", edu.StartDate, edu.EndDate)
		if edu.EndDate == "" {
			dateRange = fmt.Sprintf("%s - Present", edu.StartDate)
		}
		pdf.Cell(0, 4, dateRange)
		pdf.Ln(4)

		// Description
		if edu.Description != "" {
			pdf.SetFont("Arial", "", 10)
			pdf.SetTextColor(80, 80, 80)
			lines := s.splitText(pdf, edu.Description, 170)
			for _, line := range lines {
				pdf.Cell(0, 4, line)
				pdf.Ln(4)
			}
		}
		pdf.Ln(3)
	}
	pdf.Ln(2)
}

func (s *PDFService) addSkillsSection(pdf *gofpdf.Fpdf, skills []models.Skill) {
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(44, 62, 80)
	pdf.Cell(0, 8, "SKILLS")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(80, 80, 80)

	skillTexts := []string{}
	for _, skill := range skills {
		if skill.Level != "" {
			skillTexts = append(skillTexts, fmt.Sprintf("%s (%s)", skill.Name, skill.Level))
		} else {
			skillTexts = append(skillTexts, skill.Name)
		}
	}

	skillsText := strings.Join(skillTexts, " • ")
	lines := s.splitText(pdf, skillsText, 170)
	for _, line := range lines {
		pdf.Cell(0, 5, line)
		pdf.Ln(5)
	}
	pdf.Ln(4)
}

func (s *PDFService) addLanguagesSection(pdf *gofpdf.Fpdf, languages []string) {
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(44, 62, 80)
	pdf.Cell(0, 8, "LANGUAGES")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(80, 80, 80)

	languagesText := strings.Join(languages, " • ")
	lines := s.splitText(pdf, languagesText, 170)
	for _, line := range lines {
		pdf.Cell(0, 5, line)
		pdf.Ln(5)
	}
	pdf.Ln(4)
}

func (s *PDFService) splitText(pdf *gofpdf.Fpdf, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	var lines []string
	var currentLine string

	for _, word := range words {
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		width := pdf.GetStringWidth(testLine)
		if width <= maxWidth {
			currentLine = testLine
		} else {
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"cv-generator/internal/models"
	"cv-generator/internal/services"

	"github.com/gofiber/fiber/v2"
)

type CVHandler struct {
	pdfService *services.PDFService
}

func NewCVHandler() *CVHandler {
	return &CVHandler{
		pdfService: services.NewPDFService(),
	}
}

func (h *CVHandler) Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "CV Generator",
	})
}

func (h *CVHandler) GeneratePDF(c *fiber.Ctx) error {
	var cv models.CV

	// Parse form data
	personalInfo := models.PersonalInfo{
		FullName: c.FormValue("fullName"),
		Email:    c.FormValue("email"),
		Phone:    c.FormValue("phone"),
		Location: c.FormValue("location"),
		LinkedIn: c.FormValue("linkedin"),
		GitHub:   c.FormValue("github"),
		Website:  c.FormValue("website"),
		Summary:  c.FormValue("summary"),
	}

	// Parse education (JSON array)
	educationJSON := c.FormValue("education")
	var education []models.Education
	if educationJSON != "" {
		if err := json.Unmarshal([]byte(educationJSON), &education); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid education data"})
		}
	}

	// Parse experience (JSON array)
	experienceJSON := c.FormValue("experience")
	var experience []models.Experience
	if experienceJSON != "" {
		if err := json.Unmarshal([]byte(experienceJSON), &experience); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid experience data"})
		}
	}

	// Parse skills (JSON array)
	skillsJSON := c.FormValue("skills")
	var skills []models.Skill
	if skillsJSON != "" {
		if err := json.Unmarshal([]byte(skillsJSON), &skills); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid skills data"})
		}
	}

	// Parse languages (JSON array)
	languagesJSON := c.FormValue("languages")
	var languages []string
	if languagesJSON != "" {
		if err := json.Unmarshal([]byte(languagesJSON), &languages); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid languages data"})
		}
	}

	cv = models.CV{
		PersonalInfo: personalInfo,
		Education:    education,
		Experience:   experience,
		Skills:       skills,
		Languages:    languages,
		CreatedAt:    time.Now(),
	}

	// Generate PDF
	pdfBytes, err := h.pdfService.GenerateCV(cv)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate PDF"})
	}

	// Set headers for PDF download
	filename := fmt.Sprintf("%s_CV_%s.pdf", personalInfo.FullName, time.Now().Format("2006-01-02"))
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	return c.Send(pdfBytes)
}

func (h *CVHandler) Preview(c *fiber.Ctx) error {
	return c.Render("preview", fiber.Map{
		"Title": "CV Preview",
	})
}

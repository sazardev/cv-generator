package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("🚀 Starting PDF generation...")

	var cv models.CV

	// Parse form data
	log.Println("📝 Parsing personal info...")
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
	log.Printf("✅ Personal info parsed: %+v", personalInfo)

	// Parse education (JSON array)
	log.Println("🎓 Parsing education data...")
	educationJSON := c.FormValue("education")
	log.Printf("📋 Education JSON received: %s", educationJSON)
	var education []models.Education
	if educationJSON != "" {
		if err := json.Unmarshal([]byte(educationJSON), &education); err != nil {
			log.Printf("❌ Error parsing education: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid education data"})
		}
	}
	log.Printf("✅ Education parsed: %+v", education)

	// Parse experience (JSON array)
	log.Println("💼 Parsing experience data...")
	experienceJSON := c.FormValue("experience")
	log.Printf("📋 Experience JSON received: %s", experienceJSON)
	var experience []models.Experience
	if experienceJSON != "" {
		if err := json.Unmarshal([]byte(experienceJSON), &experience); err != nil {
			log.Printf("❌ Error parsing experience: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid experience data"})
		}
	}
	log.Printf("✅ Experience parsed: %+v", experience)

	// Parse skills (JSON array)
	log.Println("🛠️ Parsing skills data...")
	skillsJSON := c.FormValue("skills")
	log.Printf("📋 Skills JSON received: %s", skillsJSON)
	var skills []models.Skill
	if skillsJSON != "" {
		if err := json.Unmarshal([]byte(skillsJSON), &skills); err != nil {
			log.Printf("❌ Error parsing skills: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid skills data"})
		}
	}
	log.Printf("✅ Skills parsed: %+v", skills)

	// Parse languages (JSON array)
	log.Println("🗣️ Parsing languages data...")
	languagesJSON := c.FormValue("languages")
	log.Printf("📋 Languages JSON received: %s", languagesJSON)
	var languages []string
	if languagesJSON != "" {
		if err := json.Unmarshal([]byte(languagesJSON), &languages); err != nil {
			log.Printf("❌ Error parsing languages: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid languages data"})
		}
	}
	log.Printf("✅ Languages parsed: %+v", languages)

	cv = models.CV{
		PersonalInfo: personalInfo,
		Education:    education,
		Experience:   experience,
		Skills:       skills,
		Languages:    languages,
		CreatedAt:    time.Now(),
	}
	log.Printf("✅ CV struct created: %+v", cv)

	// Generate PDF
	log.Println("📄 Starting PDF generation with wkhtmltopdf...")
	pdfBytes, err := h.pdfService.GenerateCV(cv)
	if err != nil {
		log.Printf("❌ PDF generation failed: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to generate PDF: %v", err)})
	}
	log.Printf("✅ PDF generated successfully, size: %d bytes", len(pdfBytes))

	// Set headers for PDF download
	filename := fmt.Sprintf("%s_CV_%s.pdf", personalInfo.FullName, time.Now().Format("2006-01-02"))
	log.Printf("📁 Setting filename: %s", filename)
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	log.Println("🎉 PDF generation completed successfully!")
	return c.Send(pdfBytes)
}

func (h *CVHandler) Preview(c *fiber.Ctx) error {
	return c.Render("preview", fiber.Map{
		"Title": "CV Preview",
	})
}

package main

import (
	"log"
	"os"

	"cv-generator/internal/config"
	"cv-generator/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Create template engine
	engine := html.New("./web/templates", ".html")

	// Disable reload in production, enable in development
	if os.Getenv("RENDER") != "" {
		engine.Reload(false) // Production mode
	} else {
		engine.Reload(true) // Development mode
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:           "CV Generator",
		EnablePrintRoutes: true,
		Views:             engine,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Static files
	app.Static("/static", "./web/static")

	// Initialize handlers
	cvHandler := handlers.NewCVHandler()

	// Routes
	app.Get("/", cvHandler.Home)
	app.Post("/generate", cvHandler.GeneratePDF)
	app.Get("/preview", cvHandler.Preview)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"app":     "CV Generator",
			"version": "1.0.0",
			"uptime":  "running",
		})
	})

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}

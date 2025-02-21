package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes กำหนด API Routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Server is running",
		})
	})

	// เพิ่ม route อื่นๆ
}

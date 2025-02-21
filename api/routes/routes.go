package routes

import (
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/api/handler"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"gorm.io/gorm"
)

// SetupRoutes ตั้งค่า API Routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// ใช้ AssignRequestID ก่อน Log
	app.Use(middleware.AssignRequestID)
	app.Use(middleware.RequestLogger())

	// Health check
	app.Get("/health", healthcheck.New())

	// Health check: Liveness and Readiness probes
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			sqlDB, err := db.DB()
			if err != nil {
				return false
			}
			if err := sqlDB.Ping(); err != nil {
				return false
			}
			return true
		},
		ReadinessEndpoint: "/ready",
	}))

	// กำหนด API Routes เพิ่มเติม
	api := app.Group("/api")
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())
	api.Use(middleware.RateLimit())
	api.Use(middleware.GZIPCompression())

	// Breed Inquiry
	breedHandler := handler.NewBreedHandler(db)
	api.Post("/breed-inquiry", breedHandler.BreedInquiry)
}

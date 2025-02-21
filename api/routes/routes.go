package routes

import (
	"log"

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

	// Health check สำหรับ Kubernetes
	app.Get("/health", healthcheck.New()) // Health check

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

	// ตัวอย่าง Route อื่น ๆ
	// api.Post("/breed-inquiry", handler.BreedInquiry)
}


// PrintRoutes แสดงรายการ Routes ทั้งหมด
func PrintRoutes(app *fiber.App) {
	stack := app.Stack()
	log.Println("📋 Registered Routes:")
	for _, group := range stack {
		for _, route := range group {
			log.Printf("%s %s", route.Method, route.Path)
		}
	}
}

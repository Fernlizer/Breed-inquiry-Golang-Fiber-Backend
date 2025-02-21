package routes

import (
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

	// ใช้ Middleware Health Check ตามมาตรฐาน Kubernetes
	app.Get("/health", healthcheck.New())

	app.Use(healthcheck.New(healthcheck.Config{
		// 🔹 Liveness Probe → ใช้ตรวจสอบว่า Server ยังทำงานอยู่หรือไม่
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true // ถ้า Server ยังรันอยู่ ให้ return true
		},
		LivenessEndpoint: "/live", // เช็ค Liveness ที่ `/live`

		// 🔹 Readiness Probe → ใช้ตรวจสอบว่า Database พร้อมทำงานหรือไม่
		ReadinessProbe: func(c *fiber.Ctx) bool {
			sqlDB, err := db.DB()
			if err != nil {
				return false
			}
			if err := sqlDB.Ping(); err != nil {
				return false // ถ้า Database ใช้งานไม่ได้ ให้ return false
			}
			return true // ถ้าทุกอย่างพร้อมใช้งาน ให้ return true
		},
		ReadinessEndpoint: "/ready", // เช็ค Readiness ที่ `/ready`
	}))

	// Apply Global Middleware
	api := app.Group("/api")
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())
	api.Use(middleware.RateLimit())
	api.Use(middleware.GZIPCompression())

	// API Endpoints (เพิ่มในภายหลัง)
	// api.Post("/breed-inquiry", handler.BreedInquiry)
}

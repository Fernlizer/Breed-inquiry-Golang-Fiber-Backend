package main

import (
	"fmt"
	"log"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/api/routes"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/repository/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// โหลด Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// ตรวจสอบ Config
	config.ValidateConfig(cfg)

	// เลือก Database
	dbInstance, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("❌ Database setup failed:", err)
	}

	// เชื่อมต่อ Database
	db, err := dbInstance.Connect(cfg)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// ทดสอบ Database Connection
	if err := database.TestConnection(db); err != nil {
		log.Fatal("❌", err)
	}

	// สร้าง Fiber App
	app := fiber.New(fiber.Config{
		DisableStartupMessage: cfg.App.DisableStartupMessage,
		EnablePrintRoutes:     cfg.App.EnablePrintRoutes,
	})

	// กำหนด Routes
	routes.SetupRoutes(app, db)

	// รัน Server
	port := fmt.Sprintf(":%d", cfg.App.Port)
	log.Println("🚀 Server is running on port", cfg.App.Port)
	log.Fatal(app.Listen(port))
}

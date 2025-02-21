package config

import "log"

// ValidateConfig ตรวจสอบค่าที่สำคัญ
func ValidateConfig(cfg *Config) {
	if cfg.Database.User == "" || cfg.Database.Password == "" || cfg.Database.Name == "" {
		log.Fatal("❌ Database credentials are missing! Check config.yaml or ENV variables")
	}

	if cfg.App.Port == 0 {
		log.Fatal("❌ App port is not set!")
	}

	log.Println("✅ Config validation passed")
}

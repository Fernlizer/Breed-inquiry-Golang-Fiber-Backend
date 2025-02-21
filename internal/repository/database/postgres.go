package database

import (
	"fmt"
	"log"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct{}

func (p *PostgresDB) Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password,
		cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Connected to PostgreSQL:", cfg.Database.Name)
	return db, nil
}

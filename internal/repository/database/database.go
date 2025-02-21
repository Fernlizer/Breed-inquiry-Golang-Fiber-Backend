package database

import (
	"fmt"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
	"gorm.io/gorm"
)

// Database interface รองรับหลาย DB
type Database interface {
	Connect(cfg *config.Config) (*gorm.DB, error)
}

// NewDatabase เลือก Database ตาม Config
func NewDatabase(cfg *config.Config) (Database, error) {
	switch cfg.Database.Driver {
	case "postgres":
		return &PostgresDB{}, nil
	case "mysql":
		return &MySQLDB{}, nil
	case "mssql":
		return &MSSQLDB{}, nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}
}

// TestConnection ทดสอบการเชื่อมต่อ Database
func TestConnection(db *gorm.DB) error {
	var result int
	err := db.Raw("SELECT 1").Scan(&result).Error
	if err != nil || result != 1 {
		return fmt.Errorf("database connection test failed")
	}
	return nil
}

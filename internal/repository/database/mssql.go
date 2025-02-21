package database

import (
	"fmt"
	"log"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MSSQLDB struct{}

func (m *MSSQLDB) Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.Name,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Connected to MSSQL:", cfg.Database.Name)
	return db, nil
}

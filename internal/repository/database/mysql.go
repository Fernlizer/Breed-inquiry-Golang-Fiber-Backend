package database

import (
	"fmt"
	"log"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct{}

func (m *MySQLDB) Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Connected to MySQL:", cfg.Database.Name)
	return db, nil
}

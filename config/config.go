package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config โครงสร้างหลักของ Config
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Backup   BackupConfig   `mapstructure:"backup"`
}

// AppConfig กำหนดค่าของแอป
type AppConfig struct {
	Name                  string `mapstructure:"name"`
	Env                   string `mapstructure:"env"`
	Port                  int    `mapstructure:"port"`
	EnablePrefork         bool   `mapstructure:"enable_prefork"`
	DisableStartupMessage bool   `mapstructure:"disable_startup_message"`
	EnablePrintRoutes     bool   `mapstructure:"enable_print_routes"`
}

// DatabaseConfig กำหนดค่าของ Database
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

// BackupConfig สำหรับการ backup database
type BackupConfig struct {
	Enable        bool   `mapstructure:"enable"`
	PgDumpPath    string `mapstructure:"pg_dump_path"`
	RetentionDays int    `mapstructure:"retention_days"`
}

// LoadConfig โหลดค่าจากไฟล์และ ENV
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // ไม่ต้องมี ".yaml"
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")        // ค้นหาใน root directory
	viper.AddConfigPath("./config") // ค้นหาในโฟลเดอร์ config
	viper.AutomaticEnv()            // รองรับค่าจาก ENV Variables

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️ Warning: No config.yaml found, using defaults & ENV")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

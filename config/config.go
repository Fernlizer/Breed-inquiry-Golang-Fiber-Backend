package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	App struct {
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
}

// LoadConfig โหลดค่าจากไฟล์และ ENV
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // ต้องไม่มี ".yaml"
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")      // ค้นหาใน root directory
	viper.AddConfigPath("./config") // ค้นหาในโฟลเดอร์ config
	viper.AutomaticEnv()          // รองรับค่าจาก ENV Variables

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️ Warning: No config.yaml found, using defaults & ENV")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

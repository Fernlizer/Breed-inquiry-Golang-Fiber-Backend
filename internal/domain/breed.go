package domain

import (
	"time"
)

type Breed struct {
	ID        string    `gorm:"primaryKey"`
	NameTh    string    `gorm:"not null"`
	NameEn    string    `gorm:"not null"`
	ShortName string    `gorm:"size:30;not null"`
	Remark    string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

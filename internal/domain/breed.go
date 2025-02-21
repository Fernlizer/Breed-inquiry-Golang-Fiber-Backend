package domain

import (
	"time"
)

// Breed สอดคล้องกับตาราง breed ในโจทย์
type Breed struct {
	ID          string    `gorm:"type:varchar(36);primaryKey"`
	NameEn      string    `gorm:"type:varchar;not null"`
	NameTh      string    `gorm:"type:varchar;not null"`
	ShortName   string    `gorm:"type:varchar(30);not null"`
	Remark      string    `gorm:"type:varchar"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null"`
	CreatedByID string    `gorm:"type:varchar(36);not null"`
	CreatedBy   string    `gorm:"type:varchar;not null"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null"`
	UpdatedByID string    `gorm:"type:varchar(36);not null"`
	UpdatedBy   string    `gorm:"type:varchar;not null"`
}

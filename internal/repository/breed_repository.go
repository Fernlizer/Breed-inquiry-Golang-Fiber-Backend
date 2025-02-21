package repository

import (
	"log"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/domain"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/masterdata"
	"gorm.io/gorm"
)

// SeedBreeds ดึงข้อมูลจาก masterdata แล้ว insert ถ้ายังไม่มี
func SeedBreeds(db *gorm.DB) error {
	for _, breed := range masterdata.BreedMasterData {
		var existing domain.Breed
		result := db.Where("id = ?", breed.ID).First(&existing)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// ถ้าไม่พบให้ Insert
				if err := db.Create(&breed).Error; err != nil {
					log.Printf("❌ Failed to insert breed %s: %v", breed.NameEn, err)
					return err
				}
				log.Printf("✅ Inserted breed: %s", breed.NameEn)
			} else {
				// ถ้าเป็น Error อื่นให้ Return
				return result.Error
			}
		} else {
			log.Printf("🔹 Breed already exists: %s", breed.NameEn)
		}
	}
	return nil
}

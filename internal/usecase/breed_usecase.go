package usecase

import (
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/domain"
	"gorm.io/gorm"
)

// BreedUseCase ใช้สำหรับดึงข้อมูลสายพันธุ์วัว
type BreedUseCase struct {
	DB *gorm.DB
}

// NewBreedUseCase สร้างอินสแตนซ์ของ BreedUseCase
func NewBreedUseCase(db *gorm.DB) *BreedUseCase {
	return &BreedUseCase{DB: db}
}

// BreedInquiryRequest โครงสร้างข้อมูลที่รับเข้ามา
type BreedInquiryRequest struct {
	IDs        []string `json:"ids"`
	Keyword    string   `json:"keyword"`
	ShortNames []string `json:"shortNames"`
}

// SearchBreeds ค้นหาสายพันธุ์วัวจาก Database
func (u *BreedUseCase) SearchBreeds(req BreedInquiryRequest) ([]domain.Breed, error) {
	var breeds []domain.Breed
	query := u.DB.Select("id, name_th, name_en, short_name, remark").Model(&domain.Breed{})

	if len(req.IDs) > 0 {
		query = query.Where("id IN ?", req.IDs)
	}
	
	if req.Keyword != "" {
		query = query.Where("(name_th ILIKE ? OR name_en ILIKE ?)", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	
	if len(req.ShortNames) > 0 {
		query = query.Where("short_name IN ?", req.ShortNames)
	}
	
	// Execute Query
	err := query.Find(&breeds).Error
	return breeds, err
}



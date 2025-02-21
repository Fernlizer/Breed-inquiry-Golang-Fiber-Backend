package utils

import (
	"errors"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/usecase"
)

// ValidateBreedInquiryRequest ตรวจสอบค่าที่รับเข้ามา
func ValidateBreedInquiryRequest(req *usecase.BreedInquiryRequest) error {
	if req == nil {
		return errors.New("request body is required")
	}

	// ตรวจสอบว่า `ids` มีค่าหรือไม่ (ถ้าไม่ได้กำหนด `keyword` หรือ `shortNames`)
	if len(req.IDs) == 0 && req.Keyword == "" && len(req.ShortNames) == 0 {
		return errors.New("at least one of 'ids', 'keyword', or 'shortNames' must be provided")
	}

	return nil
}

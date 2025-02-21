package handler

import (
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/usecase"
	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// BreedHandler จัดการคำขอเกี่ยวกับสายพันธุ์วัว
type BreedHandler struct {
	BreedUseCase *usecase.BreedUseCase
}

// NewBreedHandler สร้างอินสแตนซ์ของ BreedHandler
func NewBreedHandler(db *gorm.DB) *BreedHandler {
	return &BreedHandler{
		BreedUseCase: usecase.NewBreedUseCase(db),
	}
}

// BreedResponse โครงสร้างข้อมูลที่ต้องการส่งกลับ
type BreedResponse struct {
	ID        string `json:"id"`
	NameTh    string `json:"nameTh"`
	NameEn    string `json:"nameEn"`
	ShortName string `json:"shortName"`
	Remark    string `json:"remark"`
}

// BreedInquiry ดึงข้อมูลสายพันธุ์วัวตามเงื่อนไขที่กำหนด
func (h *BreedHandler) BreedInquiry(c *fiber.Ctx) error {
	var req usecase.BreedInquiryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "invalid request format")
	}

	// Validate Request
	if err := utils.ValidateBreedInquiryRequest(&req); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	// Search in Database
	breeds, err := h.BreedUseCase.SearchBreeds(req)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "failed to fetch breeds")
	}

	// Map ข้อมูลให้ตรงกับรูปแบบที่โจทย์ต้องการ
	response := make([]BreedResponse, 0)
	for _, breed := range breeds {
		response = append(response, BreedResponse{
			ID:        breed.ID,
			NameTh:    breed.NameTh,
			NameEn:    breed.NameEn,
			ShortName: breed.ShortName,
			Remark:    breed.Remark,
		})
	}

	return utils.SuccessResponse(c, response)
}

package utils

import "github.com/gofiber/fiber/v2"

// SuccessResponse ใช้คืนค่า JSON ตอบกลับกรณีสำเร็จ
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

// BadRequestResponse ใช้คืนค่า JSON ตอบกลับกรณี Bad Request
func BadRequestResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": message})
}

// InternalServerErrorResponse ใช้คืนค่า JSON ตอบกลับกรณี Error ภายในระบบ
func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": message})
}

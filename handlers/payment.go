package handlers

import (
	"student-portal/database"
	"student-portal/models/students"

	"github.com/gofiber/fiber/v2"
)

// GetStudentPaymentInfo godoc
// @Summary Get student payment info
// @Description Retrieve total payable, total paid, and total due for a student
// @Tags Students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} PaymentInfo
// @Failure 404 {object} map[string]string
// @Router /students/{id}/payments [get]

func GetStudentPaymentInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	var paymentInfo struct {
		TotalPayable float64 `json:"total_payable"`
		TotalPaid    float64 `json:"total_paid"`
		TotalDue     float64 `json:"total_due"`
		PaymentStatus  string    `json:"payment_status"`
	}

	if err := database.DB.Model(&models.Student{}).
		Select("total_payable", "total_paid", "total_due", "payment_status").
		Where("id = ?", id).
		Scan(&paymentInfo).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found or payment info unavailable"})
	}

	return c.JSON(paymentInfo)
}

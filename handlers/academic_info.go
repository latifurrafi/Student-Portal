package handlers

import (
	"student-portal/database"
	"student-portal/models/students"
	"github.com/gofiber/fiber/v2"
)

func GetAcamedicInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	var academicinfo struct {
	Department      string    `json:"department"`
	Program         string    `json:"program"`
	Batch           string    `json:"batch"`
	CurrentSemester string    `json:"current_semester"`
	CreditCompleted int       `json:"credit_completed"`
	TotalCredits    int       `json:"total_credits"`
	CGPA            float32   `json:"cgpa"`
	AdmissionDate   string    `json:"admission_date"`
	}

	if err := database.DB.Model(&models.Student{}).Select("department", "program", "batch", "current_semester", "credit_completed", "total_credits", "cgpa", "admission_date").
	Where("id = ?", id).
	Scan(&academicinfo).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found or personal info unavailable."})
	}

	return c.JSON(academicinfo)
}
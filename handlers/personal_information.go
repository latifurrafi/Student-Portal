package handlers

import (
	"student-portal/database"
	"student-portal/models/students"

	"github.com/gofiber/fiber/v2"
)

func GetStudentPersonalInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	var personalinfo struct {
	Name            string    `json:"name"`
	Email           string    `json:"email" gorm:"unique"`
	Phone           string    `json:"phone"`
	DateOfBirth     string    `json:"date_of_birth"`
	Gender          string    `json:"gender"`
	BloodGroup      string    `json:"blood_group"`
	Address         string    `json:"address"`
	GuardianName     string  `json:"guardian_name"`
	GuardianPhone    string  `json:"guardian_phone"`
	GuardianRelation string  `json:"guardian_relation"`
	}

	if err := database.DB.Model(&models.Student{}).Select("name", "email", "phone", "date_of_birth", "gender", "blood_group", "address", "guardian_name", "guardian_phone", "guardian_relation").
	Where("id = ?", id).
	Scan(&personalinfo).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found or personal info unavailable"})
	}

	return c.JSON(personalinfo)
}
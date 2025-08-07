package handlers

import (
	"student-portal/database"
	"github.com/gofiber/fiber/v2"
)

func GetAvailableSemesters(c *fiber.Ctx) error {
	studentID := c.Params("student_id")

	var semesters []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}

	err := database.DB.Table("semesters").
		Select("semesters.id, semesters.name, semesters.code").
		Joins("JOIN results ON results.semester_id = semesters.id").
		Where("results.student_id = ?", studentID).
		Group("semesters.id, semesters.name, semesters.code").
		Order("semesters.id DESC").
		Scan(&semesters).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch available semesters",
		})
	}

	if len(semesters) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No semesters found for this student",
		})
	}

	// Convert to desired frontend format
	type SemesterOption struct {
		ID   string `json:"id"`
		Name string `json:"name"` // e.g., "Spring 2025, 251"
	}

	var formatted []SemesterOption
	for _, s := range semesters {
		formatted = append(formatted, SemesterOption{
			ID:   s.Code, // or s.Code if you're using string IDs
			Name: s.Name + ", " + s.Code,
		})
	}

	return c.JSON(formatted)
}

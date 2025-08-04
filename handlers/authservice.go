package handlers

import (
	"student-portal/database"
	"student-portal/models/students"
	
	"github.com/gofiber/fiber/v2"
)

func StudentLogin(c *fiber.Ctx) error {
	var req struct {
		StudentID uint 	`json:"student_id"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var student models.Student
	if err := database.DB.Where("id = ?", req.StudentID).First(&student).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	if student.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Password",
		})
	}
	 // If password matches
    return c.JSON(fiber.Map{
        "message":     "Login successful",
    })
}
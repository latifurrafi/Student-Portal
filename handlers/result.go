package handlers

import (
    "github.com/gofiber/fiber/v2"
    "student-portal/database"
    "student-portal/models"
)

func CreateResult(c *fiber.Ctx) error {
	result := new(models.Result)
	if err := c.BodyParser(result); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check if student exists
	var student models.Student
	if err := database.DB.First(&student, result.StudentID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}

	if err := database.DB.Create(&result).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save result."})
	}

	// Load the result with student information
	database.DB.Preload("Student").First(&result, result.ID)

	return c.JSON(result)
}

func GetResults(c *fiber.Ctx) error {
	var results []models.Result
	database.DB.Preload("Student").Find(&results)
	return c.JSON(results)
}
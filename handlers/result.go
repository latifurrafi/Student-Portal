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

	if err := database.DB.Create(&result).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save result."})
	}

	return c.JSON(result)
}

func GetResults(c *fiber.Ctx) error {
	var results []models.Result
	database.DB.Find(&results)
	return c.JSON(results)
}
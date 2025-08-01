package handlers

import (
	"student-portal/database"
	"student-portal/models"

	"github.com/gofiber/fiber/v2"
)

func CreateStudent(c *fiber.Ctx) error {
	student := new(models.Student)

	if err := c.BodyParser(student); err != nil{
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	result := database.DB.Create(&student)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create student."})
	}

	return c.JSON(student)
}

func GetStudents(c *fiber.Ctx) error {
	var students []models.Student
	database.DB.Preload("Results").Find(&students)
	return c.JSON(students)
}
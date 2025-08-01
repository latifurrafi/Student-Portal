package handlers

import (
    "github.com/gofiber/fiber/v2"
    "student-portal/database"
    "student-portal/models"
)


func CreateStudent(c *fiber.Ctx) error {
    var input struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    student := models.Student{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
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
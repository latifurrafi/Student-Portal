package handlers

import (
	"student-portal/database"
	models "student-portal/models/students"

	"github.com/gofiber/fiber/v2"
)

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new student with the provided information
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student Info"
// @Success 201 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [post]
func CreateStudent(c *fiber.Ctx) error {
    student := new(models.Student)
    if err := c.BodyParser(student); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid Input..."})
    }

    if err := database.DB.Create(&student).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not create student..."})
    }

    return c.Status(201).JSON(student)
}

// GetAllStudents godoc
// @Summary Get all students
// @Description Retrieve all students from the database
// @Tags Students
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {object} map[string]string
// @Router /students [get]
func GetAllStudents(c *fiber.Ctx) error {
    var students []models.Student
    if err := database.DB.Preload("Results").Find(&students).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch Students...", "detail": err.Error()})
    }
    return c.JSON(students)
}

// GetStudentByID godoc
// @Summary Get a student by ID
// @Description Retrieve a specific student by their ID
// @Tags Students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Failure 404 {object} map[string]string
// @Router /students/{id} [get]
func GetStudentsById(c *fiber.Ctx) error {
    id := c.Params("id")

    var student models.Student

    if err := database.DB.Preload("Results").First(&student, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Student not found..."})
    }

    return c.JSON(student)
}

// UpdateStudent godoc
// @Summary Update a student by ID
// @Description Update an existing student's information
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Updated student data"
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /students/{id} [put]
func UpdateStudent(c *fiber.Ctx) error {
    id := c.Params("id")
    var student models.Student

    if err := database.DB.First(&student, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Student Not Found..."})
    }

    if err := c.BodyParser(&student); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid Input..."})
    }
    
    if err := database.DB.Save(&student).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could Not Update student..."})
    } 

    return c.JSON(student)
}

// DeleteStudent godoc
// @Summary Delete a student by ID
// @Description Delete a student from the database
// @Tags Students
// @Param id path int true "Student ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /students/{id} [delete]
func DeleteStudent(c *fiber.Ctx) error {
    id := c.Params("id")

    if err := database.DB.Delete(&models.Student{}, id).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could Not Delete student..."})
    }

    return c.SendStatus(204)
}

// CreateResult godoc
// @Summary Create a new result for a student
// @Description Create a new academic result for a specific student
// @Tags Results
// @Accept json
// @Produce json
// @Param result body models.Result true "Result Info"
// @Success 201 {object} models.Result
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /results [post]
func CreateResult(c *fiber.Ctx) error {
    result := new(models.Result)
    if err := c.BodyParser(result); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input..."})
    }

    var student models.Student
    if err := database.DB.First(&student, result.StudentID).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Student Not Found..."})
    }

    if err := database.DB.Create(&result).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could Not Save result..."})
    }

    return c.Status(201).JSON(result)
}

// GetAllResults godoc
// @Summary Get all results
// @Description Retrieve all academic results from the database
// @Tags Results
// @Produce json
// @Success 200 {array} models.Result
// @Failure 500 {object} map[string]string
// @Router /results [get]
func GetAllResult(c *fiber.Ctx) error {
    var results []models.Result

    if err := database.DB.Find(&results).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch results."})
    }
    return c.JSON(results)
}

// GetResultsByStudent godoc
// @Summary Get results for a specific student
// @Description Retrieve all academic results for a specific student
// @Tags Results
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {array} models.Result
// @Failure 500 {object} map[string]string
// @Router /students/{id}/results [get]
func GetREsultsBYID(c *fiber.Ctx) error {
    StudentID := c.Params("id")
    var results []models.Result
    if err := database.DB.Where("student_id = ?", StudentID).Find(&results).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to get results..."})
    }
    return c.JSON(results)
}
package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"student-portal/cache"
	"student-portal/database"

	"github.com/gofiber/fiber/v2"
)

func GetResultsBySemester(c *fiber.Ctx) error {
	studentID := c.Params("student_id")
	semesterID := c.Params("semester_id")
	cacheKey := fmt.Sprintf("results:%s:%s", studentID, semesterID)

	// 👉 Try Redis cache first
	val, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	if err == nil { // cache hit
		var results []struct {
			CourseCode  string  `json:"course_code"`
			CourseTitle string  `json:"course_title"`
			Credit      float32 `json:"credit"`
			Grade       string  `json:"grade"`
			GradePoint  float32 `json:"grade_point"`
		}
		json.Unmarshal([]byte(val), &results)
		return c.JSON(results)
	}

	// 👉 Cache miss → query DB
	var results []struct {
		CourseCode  string  `json:"course_code"`
		CourseTitle string  `json:"course_title"`
		Credit      float32 `json:"credit"`
		Grade       string  `json:"grade"`
		GradePoint  float32 `json:"grade_point"`
	}

	err = database.DB.Table("results").
		Select("courses.code as course_code, courses.title as course_title, courses.credit, results.grade, results.grade_point").
		Joins("JOIN courses ON results.course_id = courses.id").
		Where("results.student_id = ? AND results.semester_id = ?", studentID, semesterID).
		Scan(&results).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch results",
		})
	}

	if len(results) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No results found for the given semester",
		})
	}

	// 👉 Store results in Redis for 5 minutes
	data, _ := json.Marshal(results)
	cache.Rdb.Set(cache.Ctx, cacheKey, data, 5*time.Minute)

	return c.JSON(results)
}

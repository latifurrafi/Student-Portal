package routes 

import (
	"github.com/gofiber/fiber/v2"
	"student-portal/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	student := api.Group(("/students"))
	student.Post("/", handlers.CreateStudent)
	student.Get("/", handlers.GetStudents)

    result := api.Group("/results")
    result.Post("/", handlers.CreateResult)
    result.Get("/", handlers.GetResults)
}
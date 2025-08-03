package routes 

import (
	"github.com/gofiber/fiber/v2"
	"student-portal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// api := app.Group("/api")

	app.Post("/students", handlers.CreateStudent)
	app.Get("/students", handlers.GetAllStudents)
	app.Get("/students/:id", handlers.GetStudentsById)
	app.Put("/students/:id", handlers.UpdateStudent)
	app.Delete("/students/:id", handlers.DeleteStudent)

	app.Post("/results", handlers.CreateResult)
	app.Get("/results", handlers.GetAllResult)
	app.Get("/students/:id/results", handlers.GetREsultsBYID)
}
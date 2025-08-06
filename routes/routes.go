package routes 

import (
	"github.com/gofiber/fiber/v2"
	"student-portal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Login API
	app.Post("/students/login", handlers.StudentLogin)

	// Student API
	app.Post("/students", handlers.CreateStudent)
	app.Get("/students", handlers.GetAllStudents)
	app.Get("/students/:id", handlers.GetStudentsById)
	app.Put("/students/:id", handlers.UpdateStudent)
	app.Delete("/students/:id", handlers.DeleteStudent)

	// Result API
	app.Post("/results", handlers.CreateResult)
	app.Get("/results", handlers.GetAllResult)
	app.Get("/students/:id/results", handlers.GetREsultsBYID)

	// Payment information API
	app.Get("/students/:id/payments", handlers.GetStudentPaymentInfo)

	// Personal information API
	app.Get("/students/:id/personal", handlers.GetStudentPersonalInfo)

	// Academic information API
	app.Get("/students/:id/academic", handlers.GetAcamedicInfo)
}

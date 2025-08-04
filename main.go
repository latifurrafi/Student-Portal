package main

import (
	"student-portal/database"
	"student-portal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "student-portal/docs" 
	fiberSwagger "github.com/swaggo/fiber-swagger" 
)


// @title Student Portal API
// @version 1.0
// @description This is the API documentation for the Student Portal project.
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	database.Connect()
	database.Migrate()

	routes.SetupRoutes(app)
	app.Static("/swagger", "./docs")                   
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":8000")
}

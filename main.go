package main

import (
	"log"
	"os"
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

	// PORT environment variable দিয়ে পোর্ট সেট করা
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // default port
	}

	log.Println("Starting server on port " + port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

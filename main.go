package main

import (
	"log"
	"os"
	"student-portal/cache"    // 👈 add this
	"student-portal/database"
	"student-portal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "student-portal/docs"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Student Portal API
// @version 1.0
// @description API documentation for Student Portal
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// Connect DB
	database.Connect()

	// Connect Redis
	cache.ConnectRedis()

	// Setup routes
	routes.SetupRoutes(app)

	// Swagger docs
	app.Static("/swagger", "./docs")
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Pick port from env (default 5000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("🚀 Starting server on port " + port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

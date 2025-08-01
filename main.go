package main

import (
	"github.com/gofiber/fiber/v2"
	"student-portal/database"
	"student-portal/models"
	"student-portal/routes"
)


func main() {
	app := fiber.New()

	database.Connect()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
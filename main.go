package main

import (
	"student-portal/database"
	"student-portal/routes"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())

    database.Connect()

    app.Static("/", "./public")

    routes.SetupRoutes(app)

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Api is Working.")
    })
    app.Listen(":3000")
}

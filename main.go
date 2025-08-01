package main

import (
    "github.com/gofiber/fiber/v2"
    "student-portal/database"
    "student-portal/routes"
)

func main() {
    app := fiber.New()

    database.Connect()
    routes.SetupRoutes(app)

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Api is Working.")
    })
    app.Listen(":3000")
}

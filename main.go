package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func cfg() {
	os.Getenv("PORT")
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":4000")
}

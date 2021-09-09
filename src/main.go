package main

import (
	"github.com/gofiber/fiber"
)

func index(c *fiber.Ctx) {
	c.Send("Index View")
}

func main() {
	app := fiber.New()

	app.Get("/", index)

	app.Listen(3000)
}

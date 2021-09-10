package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mentoriacompartilhada/paulushcgcj-go-v0/src/controllers"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	api := app.Group("/api")

	api.Get("/hello", controllers.Hello)

	pessoasApi := api.Group("/pessoas")

	pessoasApi.Get("", controllers.ListPessoas)
	pessoasApi.Post("", controllers.AddPessoa)

	pessoasApi.Use("/:id", controllers.HasPessoa)
	pessoasApi.Get("/:id", controllers.Getpessoa)
	pessoasApi.Put("/:id", controllers.UpdatePessoa)
	pessoasApi.Delete("/:id", controllers.DeletePessoa)

	log.Fatal(app.Listen(":3000"))
}

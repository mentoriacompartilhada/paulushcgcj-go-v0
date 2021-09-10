package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mentoriacompartilhada/paulushcgcj-go-v0/src/models"
	"github.com/mentoriacompartilhada/paulushcgcj-go-v0/src/services"
)

type Paginator struct {
	Page string `query:"page"`
	Size string `query:"size"`
}

func (p Paginator) PageInt() int {
	value, _ := strconv.ParseInt(p.Page, 10, 32)
	return int(value)
}

func (p Paginator) SizeInt() int {
	value, _ := strconv.ParseInt(p.Size, 10, 32)
	return int(value)
}

func ListPessoas(context *fiber.Ctx) error {

	paginator := new(Paginator)

	if err := context.QueryParser(paginator); err != nil {
		context.SendStatus(500)
		return context.SendString(err.Error())
	}

	return context.JSON(services.ListPessoas(paginator.PageInt(), paginator.SizeInt()))
}

func AddPessoa(context *fiber.Ctx) error {
	var pessoa models.Pessoa
	if err := context.BodyParser(&pessoa); err != nil {
		context.SendStatus(400)
		return context.SendString(err.Error())
	}
	pessoaId := services.AddPessoa(pessoa)
	context.Append("Location", fmt.Sprintf("/api/pessoas/%d", pessoaId))
	context.SendStatus(201)
	return context.SendString("")
}

func Getpessoa(context *fiber.Ctx) error {
	pessoaId, _ := context.ParamsInt("id")
	return context.JSON(services.Getpessoa(pessoaId))
}

func UpdatePessoa(context *fiber.Ctx) error {
	pessoaId, _ := context.ParamsInt("id")
	var pessoa models.Pessoa
	if err := context.BodyParser(&pessoa); err != nil {
		context.SendStatus(400)
		return context.SendString(err.Error())
	}
	services.UpdatePessoa(pessoaId, pessoa)
	context.SendStatus(202)
	return context.SendString("")
}

func DeletePessoa(context *fiber.Ctx) error {
	pessoaId, _ := context.ParamsInt("id")
	services.DeletePessoa(pessoaId)
	context.SendStatus(204)
	return context.SendString("")
}

func HasPessoa(context *fiber.Ctx) error {
	pessoaId, error := context.ParamsInt("id")

	if error != nil {
		context.SendStatus(404)
		return context.SendString(error.Error())
	}

	if services.HasPessoa(pessoaId) {
		return context.Next()
	}
	return context.SendStatus(404)
}

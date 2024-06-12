package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				ctx.Status(fiber.StatusBadRequest)
				return ctx.SendString("error : " + err.Error())
			},
		},
	)

	app.Get("/error", func(ctx *fiber.Ctx) error {
		return errors.New("ups error")
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello")
	})

	app.Listen("localhost:8080")

}

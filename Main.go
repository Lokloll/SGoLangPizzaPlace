package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./src/")
	app.Get("/formin", Formin)
	app.Post("/", func(ctx *fiber.Ctx) error {
		file, _ := ctx.FormFile("src/index.html")

		return ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})

	app.Get("/shutdown", func(ctx *fiber.Ctx) error {
		return app.Shutdown()
	})
	app.Listen(":8880")
}

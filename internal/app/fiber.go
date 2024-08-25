package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
)

var _ = di.Provide(NewFiber)

func NewFiber() *fiber.App {
	app := fiber.New()
	app.Static("/", "./public")
	return app
}

package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/config"
)

func Start(app *fiber.App, cfg *config.Config) error {
	return app.Listen(cfg.Address)
}

func Stop(app *fiber.App) error {
	return app.Shutdown()
}

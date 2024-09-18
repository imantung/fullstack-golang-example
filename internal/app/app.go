package app

import (
	"context"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/config"
	"github.com/labstack/echo/v4"
)

func Start(e *echo.Echo, cfg *config.Config) error {
	return e.Start(cfg.Address)
}

func Stop(e *echo.Echo) error {
	ctx := context.Background()
	return e.Shutdown(ctx)
}

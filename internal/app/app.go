package app

import (
	"context"
	"net/http"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/config"
	"github.com/labstack/echo/v4"
)

func Start(e *echo.Echo, cfg *config.Config) error {
	return e.Start(cfg.Address)
}

func Stop(server *http.Server) error {
	ctx := context.Background()
	return server.Shutdown(ctx)
}

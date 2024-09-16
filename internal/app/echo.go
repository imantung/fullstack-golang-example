package app

import (
	"net/http"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/config"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

var _ = di.Provide(NewEcho)

func NewEcho(cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.Renderer = NewTemplateRegistry()

	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", nil)
	})
	return e
}

// https://echo.labstack.com/docs/templates
// https://medium.com/free-code-camp/how-to-setup-a-nested-html-template-in-the-go-echo-web-framework-670f16244bb4

package app

import (
	"github.com/imantung/fullstack-golang-example/internal/app/controller"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

var _ = di.Provide(NewEcho)

func NewEcho(
	webCntrl controller.WebController,
) *echo.Echo {

	e := echo.New()
	e.Renderer = NewTemplateRegistry()

	e.Static("/dist", "web/dist")
	e.GET("/", webCntrl.HomePage)
	e.GET("/about", webCntrl.AboutPage)

	return e
}

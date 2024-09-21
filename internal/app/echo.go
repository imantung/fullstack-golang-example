package app

import (
	"net/http"

	"github.com/imantung/fullstack-golang-example/internal/app/controller"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var _ = di.Provide(NewEcho)

func NewEcho(
	webCntrl controller.WebController,
) *echo.Echo {

	e := echo.New()
	e.Renderer = NewTemplateRegistry()
	e.HTTPErrorHandler = HTTPErrorHandler

	e.Use(middleware.Recover())

	e.Static("/dist", "web/dist")
	e.GET("/", webCntrl.HomePage)
	e.GET("/about", webCntrl.AboutPage)
	e.GET("/child", webCntrl.ChildPage)

	return e
}

func HTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	c.Render(http.StatusOK, "error.html", he)
}

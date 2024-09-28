package app

import (
	"net/http"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

var _ = di.Provide(NewEcho)

func NewEcho(router Router, templateRegistry *TemplateRegistry) *echo.Echo {

	e := echo.New()
	e.Renderer = templateRegistry
	e.HTTPErrorHandler = HTTPErrorHandler

	router.SetRoute(e)

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

package controller

import (
	"net/http"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	WebController interface {
		HomePage(c echo.Context) error
		AboutPage(c echo.Context) error
	}
	WebControllerImpl struct{}
)

var _ = di.Provide(NewWebController)

func NewWebController() WebController {
	return &WebControllerImpl{}
}

func (*WebControllerImpl) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func (*WebControllerImpl) AboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", nil)
}
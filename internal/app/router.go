package app

import (
	"path/filepath"

	"github.com/imantung/fullstack-golang-example/internal/app/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

type Router struct {
	dig.In
	WebCntrl controller.WebController
}

func (r *Router) SetRoute(e *echo.Echo) {
	e.Use(middleware.Recover())

	e.Static("/dist", filepath.Join(RootPath, DistPath))
	e.GET("/", r.WebCntrl.HomePage)
	e.GET("/about", r.WebCntrl.AboutPage)
	e.GET("/child", r.WebCntrl.ChildPage)
	e.GET("/grand-child", r.WebCntrl.GrandChildPage)
}

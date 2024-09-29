package app

import (
	"embed"

	"github.com/imantung/fullstack-golang-example/internal/app/controller"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Router struct {
	dig.In
	WebCntrl controller.WebController
}

var (
	StaticPath = "view/static"

	//go:embed all:view/static/*
	StaticFS embed.FS
)

func (r *Router) SetRoute(e *echo.Echo) {
	// e.Use(middleware.Recover())

	e.StaticFS("/static", echo.MustSubFS(StaticFS, StaticPath))
	e.GET("/", r.WebCntrl.HomePage)
	e.GET("/about", r.WebCntrl.AboutPage)
	e.GET("/child", r.WebCntrl.ChildPage)
	e.GET("/grand-child", r.WebCntrl.GrandChildPage)
}

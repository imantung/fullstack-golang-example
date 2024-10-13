package app

import (
	"embed"
	"html/template"
	"io"
	"net/http"

	"github.com/imantung/dirtmpl"
	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	TemplateRegistry struct {
		Templates map[string]*template.Template
	}
)

var (
	Root = "view/pages"

	//go:embed all:view/pages/*
	TemplatesFS embed.FS
)

var _ = di.Provide(NewTemplateRegistry)

func NewTemplateRegistry() *TemplateRegistry {
	templates, _ := dirtmpl.HTMLTemplatesFS(TemplatesFS, Root)

	return &TemplateRegistry{Templates: templates}
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if template, ok := t.Templates[name]; ok {
		return template.Execute(w, data)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, "missing template: "+name)
}

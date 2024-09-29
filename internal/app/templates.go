package app

import (
	"embed"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/imantung/fullstack-golang-example/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

type (
	TemplateRegistry struct {
		Templates map[string]*template.Template
	}
)

var (
	TemplatePath = "view/templates"

	//go:embed all:view/templates/*
	TemplatesFS embed.FS
)

var _ = di.Provide(NewTemplateRegistry)

func NewTemplateRegistry() *TemplateRegistry {
	var m map[string][]string = make(map[string][]string)
	WalkTemplate(m, TemplatePath, []string{})

	templates := map[string]*template.Template{}

	for k, v := range m {
		templates[k] = template.Must(template.ParseFS(TemplatesFS, v...))
	}

	return &TemplateRegistry{Templates: templates}
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if template, ok := t.Templates[name]; ok {
		return template.Execute(w, data)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, "missing template: "+name)
}

func WalkTemplate(m map[string][]string, parent string, list []string) {
	entries, _ := TemplatesFS.ReadDir(parent)
	for _, entry := range entries {
		filename := entry.Name()
		fullPath := filepath.Join(parent, filename)

		if strings.HasPrefix(filename, "_") {
			list = append(list, fullPath)
		} else {
			if entry.IsDir() {
				WalkTemplate(m, fullPath, list)
			} else {
				list2 := make([]string, len(list))
				copy(list2, list)

				key := fullPath[len(TemplatePath)+1:]
				m[key] = append(list2, fullPath)
			}
		}
	}
}

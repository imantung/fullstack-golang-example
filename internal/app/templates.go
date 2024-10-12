package app

import (
	"embed"
	"fmt"
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
	Root = "view/pages"

	//go:embed all:view/pages/*
	TemplatesFS embed.FS
)

var _ = di.Provide(NewTemplateRegistry)

func NewTemplateRegistry() *TemplateRegistry {
	var m map[string][]string = make(map[string][]string)
	WalkTemplates(m, Root, []string{})

	templates := map[string]*template.Template{}

	for k, v := range m {
		if k == "home.html" {
			fmt.Println(k, v)
			templates[k] = template.Must(template.ParseFS(TemplatesFS, v...))
		}
	}

	return &TemplateRegistry{Templates: templates}
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if template, ok := t.Templates[name]; ok {
		return template.Execute(w, data)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, "missing template: "+name)
}

func WalkTemplates(m map[string][]string, parent string, bases []string) {
	entries, _ := TemplatesFS.ReadDir(parent)
	var comps []string
	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(parent, name)

		if strings.HasPrefix(name, "_") {
			if entry.IsDir() {
				entries, _ := TemplatesFS.ReadDir(path)
				for _, entry := range entries {
					if !entry.IsDir() {
						comps = append(comps, filepath.Join(parent, name, entry.Name()))
					}
				}
			} else {
				bases = append(bases, path)
			}
		} else {
			if entry.IsDir() {
				WalkTemplates(m, path, bases)
			} else {
				includes := make([]string, len(bases))
				copy(includes, bases)
				includes = append(includes, comps...)
				key := path[len(Root)+1:]
				m[key] = append(includes, path)
			}
		}
	}
}

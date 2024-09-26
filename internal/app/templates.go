package app

import (
	"embed"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type (
	TemplateRegistry struct {
		Templates map[string]*template.Template
	}
)

var (
	RootPath     = "internal/app/"
	TemplatePath = "view/src/templates/"
	DistPath     = "view/dist"

	//go:embed all:view/src/templates/*
	TemplatesFS embed.FS
)

func NewTemplateRegistry() *TemplateRegistry {
	var dict map[string][]string = make(map[string][]string)
	WalkTemplate(dict, "", []string{})

	templates := map[string]*template.Template{}
	for k, v := range dict {
		templates[k] = template.Must(template.ParseFS(TemplatesFS, v...))
	}
	return &TemplateRegistry{Templates: templates}
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if template, ok := t.Templates[name]; ok {
		return template.Execute(w, data)
	}
	return t.Templates["error-404.html"].Execute(w, data)
}

func WalkTemplate(dict map[string][]string, parent string, list []string) {
	parent = filepath.Join(TemplatePath, parent)
	files, _ := os.ReadDir(filepath.Join(RootPath, parent))

	for _, file := range files {
		fullPath := filepath.Join(parent, file.Name())
		relPath := fullPath[len(TemplatePath):]
		if file.Name() == "_base.html" {
			list = append(list, fullPath)
			continue
		}
		if file.IsDir() {
			WalkTemplate(dict, relPath, list)
		} else {
			dict[relPath] = append(list, fullPath)
		}
	}
}

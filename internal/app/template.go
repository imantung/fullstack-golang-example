package app

import (
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

func NewTemplateRegistry() *TemplateRegistry {
	var dict map[string][]string = make(map[string][]string)

	src := "web/src/templates/"

	list := []string{}
	walkFile(dict, src, "", list)

	templates := map[string]*template.Template{}
	for k, v := range dict {
		templates[k] = template.Must(template.ParseFiles(v...))
	}
	return &TemplateRegistry{Templates: templates}
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if template, ok := t.Templates[name]; ok {
		return template.Execute(w, data)
	}
	return t.Templates["error-404.html"].Execute(w, data)
}

func walkFile(dict map[string][]string, src, parent string, list []string) {
	if src == "" {
		parent = src
	} else {
		parent = filepath.Join(src, parent)
	}

	files, _ := os.ReadDir(parent)

	for _, file := range files {
		fullPath := filepath.Join(parent, file.Name())
		relPath := fullPath[len(src):]

		if file.Name() == "_base.html" {
			list = append(list, fullPath)
			continue
		}

		if file.IsDir() {
			walkFile(dict, src, relPath, list)
		} else {
			dict[relPath] = append(list, fullPath)
		}
	}
}

package template

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type EchoConfig struct {
	Function template.FuncMap
}

type EchoTemplate struct {
	templates map[string]*template.Template
}

func (t *EchoTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].ExecuteTemplate(w, name, data)
}

func EchoRender(config EchoConfig) *EchoTemplate {
	includeDir := "templates/includes"
	layoutDir := "templates/layouts"

	ts := map[string]*template.Template{}

	includes, layouts := getTemplateFilePath(includeDir, layoutDir)

	for _, layout := range layouts {
		name := layout
		files := append(includes, layout)
		var tmpl *template.Template
		if config.Function != nil {
			tmpl = template.Must(template.New(name).Funcs(defaultFuncs()).Funcs(config.Function).ParseFiles(files...))
		} else {
			tmpl = template.Must(template.New(name).Funcs(defaultFuncs()).ParseFiles(files...))
		}
		ts[name] = tmpl
	}

	return &EchoTemplate{
		templates: ts,
	}
}

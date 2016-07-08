package template

import (
	common "github.com/hixi-hyi/gotch/waf/common/template"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Config struct {
	Function template.FuncMap
}

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].ExecuteTemplate(w, name, data)
}

func Render(config Config) *Template {
	includeDir := "templates/includes"
	layoutDir := "templates/layouts"

	ts := map[string]*template.Template{}

	includes, layouts := common.GetTemplateFilePath(includeDir, layoutDir)

	for _, layout := range layouts {
		name := layout
		files := append(includes, layout)
		var tmpl *template.Template
		if config.Function != nil {
			tmpl = template.Must(template.New(name).Funcs(common.DefaultFuncs()).Funcs(config.Function).ParseFiles(files...))
		} else {
			tmpl = template.Must(template.New(name).Funcs(common.DefaultFuncs()).ParseFiles(files...))
		}
		ts[name] = tmpl
	}

	return &Template{
		templates: ts,
	}
}

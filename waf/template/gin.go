package template

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"html/template"
)

type GinConfig struct {
	Function template.FuncMap
}

func GinRender(config GinConfig) multitemplate.Render {
	includeDir := "templates/includes"
	layoutDir := "templates/layouts"

	r := multitemplate.New()
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
		r.Add(name, tmpl)
	}
	return r
}

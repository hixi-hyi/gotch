package template

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
	common "github.com/hixi-hyi/gotch/waf/common/template"
	"html/template"
)

type Config struct {
	Function template.FuncMap
}

func Render(config Config) multitemplate.Render {
	includeDir := "templates/includes"
	layoutDir := "templates/layouts"

	r := multitemplate.New()
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
		r.Add(name, tmpl)
	}
	return r
}

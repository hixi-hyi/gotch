package template

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"html/template"
	"os"
	"path/filepath"
)

type GinConfig struct {
	Function template.FuncMap
}

func deepRead(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		*files = append(*files, path)
		return nil
	}
}

func GinRender(config GinConfig) multitemplate.Render {
	includeDir := "templates/includes"
	layoutDir := "templates/layouts"

	r := multitemplate.New()

	var layouts, includes []string
	filepath.Walk(layoutDir, deepRead(&layouts))
	filepath.Walk(includeDir, deepRead(&includes))

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

// dummy
func defaultFuncs() template.FuncMap {
	return template.FuncMap{}
}

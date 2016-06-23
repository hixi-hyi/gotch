package template

import (
	"errors"
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

func defaultFuncs() template.FuncMap {
	return template.FuncMap{
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	}
}

package template

import (
	"errors"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
)

func deepRead(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		*files = append(*files, path)
		return nil
	}
}

func GetTemplateFilePath(includeDir, layoutDir string) ([]string, []string) {
	var layouts, includes []string
	filepath.Walk(layoutDir, deepRead(&layouts))
	filepath.Walk(includeDir, deepRead(&includes))
	return includes, layouts
}

func DefaultFuncs() template.FuncMap {
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
		"lineBreak": lineBreak,
	}
}

func lineBreak(s string) template.HTML {
	rep := regexp.MustCompile(`[\r\n]`)
	return template.HTML(rep.ReplaceAllString(s, "<br>"))
}

package template

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func parse(path string, _ fs.FileInfo, err error) error {
	if strings.Contains(path, ".html") {
		tmplBytes, err := os.ReadFile(path)

		// Failed to read file
		if err != nil {
			return err
		}

		_, err = template.New(path).Funcs(sprig.FuncMap()).Parse(string(tmplBytes))
		if err != nil {
			return err
		}
	}
	return err
}

func ParseTemplates() (*template.Template, error) {
	template := template.New("").Funcs(sprig.FuncMap())

	err := filepath.Walk("templates", parse)

	if err != nil {
		return nil, err
	}

	return template, nil
}

func GetTemplate() {

}

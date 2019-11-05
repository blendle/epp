package epp

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
)

const DefaultPartialsPath = "config/deploy/partials"

// Parse parses the input and returns the output
func Parse(input []byte, partialPath string) ([]byte, error) {
	var writer bytes.Buffer
	t := template.New("test")

	funcMap := template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			buf := bytes.NewBuffer(nil)
			if err := t.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},

		"required": func(warn string, val interface{}) (interface{}, error) {
			if val == nil {
				return val, fmt.Errorf(warn)
			} else if _, ok := val.(string); ok {
				if val == "" {
					return val, fmt.Errorf(warn)
				}
			}
			return val, nil
		},
	}

	err := loadTemplates(t, partialPath)
	if err != nil {
		return nil, err
	}

	t, tplErr := t.Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(string(input))
	if tplErr != nil {
		return nil, err
	}

	err = t.Execute(&writer, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func loadTemplates(t *template.Template, partialPath string) error {
	if partialPath == DefaultPartialsPath {
		if _, err := os.Stat(partialPath); os.IsNotExist(err) {
			return nil
		}
	}

	if partialPath == "" {
		return nil
	}

	var templatePaths []string
	err := filepath.Walk(partialPath, func(path string, f os.FileInfo, err error) error {
		if info, e := os.Stat(path); e == nil && info.IsDir() {
			return nil
		}
		templatePaths = append(templatePaths, path)

		return nil
	})

	if err != nil {
		return err
	}

	_, err = t.ParseFiles(templatePaths...)
	return err
}

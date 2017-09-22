package epp

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
)

// Parse parses the input and returns the output
func Parse(input []byte) ([]byte, error) {
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
	}

	tpl, err := t.Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(string(input))
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(&writer, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

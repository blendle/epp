package epp

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
)

// Parse parses the input and returns the output
func Parse(input []byte) ([]byte, error) {
	var writer bytes.Buffer

	tpl, err := template.New("test").Funcs(sprig.TxtFuncMap()).Parse(string(input))
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(&writer, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

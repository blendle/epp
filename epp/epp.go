package epp

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/flosch/pongo2"
)

func init() {
	pongo2.RegisterFilter("b64enc", filterBase64Encode)
}

// Parse parses the input and returns the output
func Parse(input []byte) ([]byte, error) {
	tpl, err := pongo2.FromString(string(input))
	if err != nil {
		return nil, err
	}

	context := environToContext()
	return tpl.ExecuteBytes(context)
}

func environToContext() pongo2.Context {
	ctx := pongo2.Context{}

	for _, env := range os.Environ() {
		variable := strings.SplitN(env, "=", 2)
		key, value := variable[0], variable[1]

		ctx[key] = value
	}

	return ctx
}

func filterBase64Encode(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(base64.StdEncoding.EncodeToString([]byte(in.String()))), nil
}

package epp

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	os.Setenv("SPLIT_TEST", "value=with=equal=sings")
	os.Setenv("KUBERNETES_ADDRESS", "https://192.168.99.100")
	os.Setenv("FALSY", "")
}

func TestEnvVariables(t *testing.T) {
	tpl := []byte(`{{ env "SPLIT_TEST" }}: {{ env "KUBERNETES_ADDRESS" }}`)
	expected := fmt.Sprintf("%s: %s", os.Getenv("SPLIT_TEST"), os.Getenv("KUBERNETES_ADDRESS"))

	res, err := Parse(tpl)

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestEnvIf(t *testing.T) {
	tpl := []byte(`
{{- if env "$FALSY" }}
I shouldn't appear
{{- end }}
I should!
`)
	expected := `
I should!
`

	res, err := Parse(tpl)

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestInclude(t *testing.T) {
	tpl := []byte(`{{ define "worldtpl" }}world{{- end }}
hello {{ include "worldtpl" . | upper }}`)
	expected := `
hello WORLD`

	res, err := Parse(tpl)

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

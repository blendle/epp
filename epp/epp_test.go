package epp

import (
	"fmt"
	"os"
	"strings"
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

	res, err := Parse(tpl, "")

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

	res, err := Parse(tpl, "")

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

	res, err := Parse(tpl, "")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestRequired(t *testing.T) {
	tpl := []byte(`{{ required "undefined" "hello" }}`)
	expected := `hello`

	res, err := Parse(tpl, "")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestRequired_Undefined(t *testing.T) {
	tpl := []byte(`{{ required "undefined" .Undefined }}`)
	_, err := Parse(tpl, "")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "error calling required: undefined") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStringList(t *testing.T) {
	tpl := []byte(`{{ splitList " " "hello world" | last }}`)
	expected := `world`

	res, err := Parse(tpl, "")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestPartialDir(t *testing.T) {
	tpl := []byte(`{{ include "hello_partial" . }},{{ include "world_partial" . }}`)
	expected := `Hello,world`

	res, err := Parse(tpl, "../resources")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

func TestPartialFunctions(t *testing.T) {
	tpl := []byte(`{{ include "recursive_partial" . }}`)
	expected := `Hello,world`

	res, err := Parse(tpl, "../resources")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if string(res) != expected {
		t.Errorf("bad expansion: expected '%s', got '%s'", expected, res)
	}
}

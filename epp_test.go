package main

import (
	"testing"
)

func TestSingleFileInput(t *testing.T) {
	result, err := readInput("resources/test_template.yml")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	expected := `---
test: bla
`

	if string(result) != expected {
		t.Errorf("unexpected content '%s' but expected: '%s'", result, expected)
	}
}

func TestWildcardFile(t *testing.T) {
	result, err := readInput("resources/*.yml")

	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	expected := `---
test: bla
---
template2: success
`

	if string(result) != expected {
		t.Errorf("unexpected content '%s' but expected: '%s'", result, expected)
	}
}

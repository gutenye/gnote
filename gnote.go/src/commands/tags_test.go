package commands

import (
	"testing"
)

func TestExtractTagsFromText(t *testing.T) {
	text := `
		hello world
		foo *bar*
		link step
		*baz* and 
	`
	expected :=
		"bar\t/a.gnote\t/*bar*\n" +
			"baz\t/a.gnote\t/*baz*"
	result := extractTagsFromText(text, "/a.gnote", "*", createPattern("*"))
	if result != expected {
		t.Error("Failed")
	}
}

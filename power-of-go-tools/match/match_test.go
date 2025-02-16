package match_test

import (
	"bytes"
	"testing"

	"github.com/netletic/match"
)

func TestMatchString(t *testing.T) {
	t.Parallel()
	text := `this won't match
this line is MAGIC

MAGIC is also on this line
but not this one
and MAGIC is here too
`
	input := bytes.NewBufferString(text)
	output := bytes.NewBufferString("")
	m, err := match.NewMatcher(
		match.WithInput(input),
		match.WithOutput(output),
		match.WithSearchTextFromArgs([]string{"MAGIC"}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := `this line is MAGIC
MAGIC is also on this line
and MAGIC is here too
`
	m.PrintMatchingLines()
	got := output.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

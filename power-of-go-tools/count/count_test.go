package count_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/netletic/count"

	"github.com/rogpeppe/go-internal/testscript"
)

// maps the "exec" commands in the *.txtar files within the testdata/script
// dir to func count.Main from count.go
func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"lines": count.MainLines,
		"words": count.MainWords,
		"count": count.Main,
	}))
}

// executes all *.txtar files within the testdata/script dir as tests
func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestLinesCountLinesInInput(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(count.WithInput(input))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(count.WithInputFromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	args := []string{}
	input := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(input),
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLinesCountWordsInInput(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("1\nfoo bar baz\nqux quux")
	c, err := count.NewCounter(count.WithInput(input))
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

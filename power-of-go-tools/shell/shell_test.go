package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/shell"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"shell": shell.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	input := ""
	_, err := shell.CmdFromString(input)
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}
func TestCmdFromString_CreatesExpectedCmd(t *testing.T) {
	t.Parallel()
	input := "/bin/ls -l main.go"
	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"/bin/ls", "-l", "main.go"}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNewSession_CreatesExpectedSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Stdin:      os.Stdin,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		Transcript: io.Discard,
		DryRun:     false,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun_ProducesExpectedOutput(t *testing.T) {
	t.Parallel()
	stdin := strings.NewReader("echo hello\n\n")
	stdout := new(bytes.Buffer)
	session := shell.NewSession(stdin, stdout, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nBe seeing you!\n"
	got := stdout.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRun_ProducesExpectedTranscript(t *testing.T) {
	t.Parallel()
	stdin := strings.NewReader("echo hello\n\n")
	transcript := new(bytes.Buffer)
	session := shell.NewSession(stdin, io.Discard, io.Discard)
	session.DryRun = true
	session.Transcript = transcript
	session.Run()
	want := "> echo hello\necho hello\n> \n> \nBe seeing you!\n"
	got := transcript.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

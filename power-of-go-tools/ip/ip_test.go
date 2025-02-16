package ip_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/ip"
)

func TestExtractIP(t *testing.T) {
	input := `192.168.1.10 - - [28/Jan/2025:14:22:11 +0000] "GET /index.html HTTP/1.1" 200 1024`
	want := "192.168.1.10"
	got, err := ip.ExtractIP(input)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestExtractIP_ReturnsErrorIfExtractFails(t *testing.T) {
	input := `192.foo.1.10 - - [28/Jan/2025:14:22:11 +0000] "GET /index.html HTTP/1.1" 200 1024`
	_, err := ip.ExtractIP(input)
	if err == nil {
		t.Fatal("want error when correct IP can't be extracted from the input line")
	}
}

func TestMain(t *testing.T) {
	want := map[string]int{
		"192.168.1.10":  5,
		"203.0.113.42":  4,
		"198.51.100.25": 4,
		"192.168.1.20":  1,
		"192.168.1.30":  1,
	}
	got := ip.Main()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

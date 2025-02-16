package older_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/older"
)

func TestFiles_OlderThan30Days(t *testing.T) {
	new := time.Now()
	old := new.Add(-time.Minute)
	fsys := fstest.MapFS{
		"new.txt":            {ModTime: new},
		"old.txt":            {ModTime: old},
		"subfolder/new.txt":  {ModTime: new},
		"subfolder/old.txt":  {ModTime: old},
		"subfolder2/new.txt": {ModTime: new},
		"subfolder2/old.txt": {ModTime: old},
	}
	want := []string{
		"old.txt",
		"subfolder/old.txt",
		"subfolder2/old.txt",
	}
	got := older.Than(fsys, time.Second*10)
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

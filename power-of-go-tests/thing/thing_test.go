package thing_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/thing"
)

func TestNewThing(t *testing.T) {
	t.Parallel()
	want := &thing.Thing{
		X: 1,
		Y: 2,
		Z: 3,
	}
	got, err := thing.NewThing(1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		cmp.Diff(want, got)
	}
}

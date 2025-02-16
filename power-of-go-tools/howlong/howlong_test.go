package howlong_test

import (
	"testing"
	"time"

	"github.com/netletic/howlong"
)

const epsilon = 200 * time.Millisecond

func TestExecMeasuresElapsedTime(t *testing.T) {
	t.Parallel()
	want := 100 * time.Millisecond
	got, err := howlong.Run("sleep", "0.1")
	if err != nil {
		t.Fatal(err)
	}
	epsilon := (want - got).Abs()
	if epsilon > 300*time.Millisecond {
		t.Errorf("want %s, got %s (not close enough)", want, got)
	}
}

package game_test

import (
	"game"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListItems_GivesCorrectResultFor(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name  string
		input []string
		want  string
	}

	tcs := []testCase{
		{
			name: "three items",
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
		{
			name: "two items",
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
		{
			name: "one item",
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		{
			name:  "no items",
			input: []string{},
			want:  "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := game.ListItems(tc.input)
			if !cmp.Equal(tc.want, got) {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}

package adventure_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/adventure"
)

func TestFromJSON_CorrectlyParsesAdventureFromJSONFile(t *testing.T) {
	got, err := adventure.FromJSON("testdata/adventure.json")
	if err != nil {
		t.Fatal(err)
	}

	want := adventure.Adventure{
		"intro": {
			Title: "The Beginning",
			Story: []string{"Once upon a time..."},
			Options: []adventure.Option{
				{
					Text:    "Go to the next chapter",
					Chapter: "next_chapter",
				},
			},
		},
		"next_chapter": {
			Title:   "The Next Chapter",
			Story:   []string{"The story continues..."},
			Options: []adventure.Option{},
		},
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

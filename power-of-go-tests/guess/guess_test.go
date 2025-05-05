package guess_test

import (
	"testing"

	"github.com/netletic/guess"
)

func FuzzGuess(f *testing.F) {
	// f.Add(21)
	f.Fuzz(func(t *testing.T, input int) {
		guess.Guess(input)
	})
}

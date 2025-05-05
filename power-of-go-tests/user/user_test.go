package user_test

import (
	"testing"

	"github.com/netletic/user"
)

func TestGreetingReturnsCorrectGreetingForLanguage(t *testing.T) {
	t.Parallel()
	u := user.New("Fake User")
	u.Language = "French"
	want := "Bonjour"
	got := u.Greeting()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

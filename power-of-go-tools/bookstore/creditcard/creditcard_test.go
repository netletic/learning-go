package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "123456789"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewInvalidReturnsError(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Error("want error for invalid card number, got nil")
	}
}

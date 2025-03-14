package kv_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/kv"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"kv": kv.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore(t.TempDir() + "dummy path")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("nonexistent-key")
	if ok {
		t.Fatal("unexpected ok for nonexistent key")
	}
}

func TestGetReturnsValueAndOkIfKeyExists(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore(t.TempDir() + "dummy path")
	if err != nil {
		t.Fatal(err)
	}
	want := "value"
	s.Set("key", want)
	got, ok := s.Get("key")
	if !ok {
		t.Fatal("wanted ok for existing key")
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSetOverwritesExistingValue(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore(t.TempDir() + "dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "original")
	want := "updated"
	s.Set("key", want)
	got, ok := s.Get("key")
	if !ok {
		t.Fatal("key not found")
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestAllReturnsExpectedMap(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	want := map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
	}
	got := s.All()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSaveSavesDataPersistently(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}
	s2, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := s2.Get("A"); v != "1" {
		t.Fatalf("want A=1, got A=%s", v)
	}
	if v, _ := s2.Get("B"); v != "2" {
		t.Fatalf("want B=2, got B=%s", v)
	}
	if v, _ := s2.Get("C"); v != "3" {
		t.Fatalf("want C=3, got C=%s", v)
	}
}

func TestOpenStore_ErrorsWhenPathUnreadable(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/unreadable.store"
	if _, err := os.Create(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0o000); err != nil {
		t.Fatal(err)
	}
	_, err := kv.OpenStore(path)
	if err == nil {
		t.Fatal("want error if kv storefile is unreadable")
	}
}

func TestOpenStore_ReturnsErrorOnInvalidData(t *testing.T) {
	t.Parallel()
	_, err := kv.OpenStore("testdata/invalid_data.store")
	if err == nil {
		t.Fatal("want error if kv storefile contains invalid data")
	}
}

func TestSaveErrorsWhenPathUnwriteable(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("bogus/unwritable.store")
	if err != nil {
		t.Fatal(err)
	}
	err = s.Save()
	if err == nil {
		t.Fatal("want error if kv storefile is unwritable")
	}
}

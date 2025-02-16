package quiz_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/quiz"
)

func TestFromCSV_CorrectlyReadsProblemsFromCSVFile(t *testing.T) {
	t.Parallel()
	want := []quiz.Problem{
		{Question: "5+5", Answer: "10"},
		{Question: "7+3", Answer: "10"},
		{Question: "1+1", Answer: "2"},
		{Question: "8+3", Answer: "11"},
		{Question: "1+2", Answer: "3"},
		{Question: "8+6", Answer: "14"},
		{Question: "3+1", Answer: "4"},
		{Question: "1+4", Answer: "5"},
		{Question: "5+1", Answer: "6"},
		{Question: "2+3", Answer: "5"},
		{Question: "3+3", Answer: "6"},
		{Question: "2+4", Answer: "6"},
		{Question: "5+2", Answer: "7"},
	}
	got, err := quiz.FromCSV("testdata/problems.csv")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFromCSV_ReturnsErrorForUnreadableCSVFile(t *testing.T) {
	t.Parallel()
	_, err := quiz.FromCSV(t.TempDir() + "bogus.csv")
	if err == nil {
		t.Fatal("want error if CSV file is unreadable.")
	}
}

func TestFromCSV_ReturnsErrorGivenInvalidCSVFile(t *testing.T) {
	t.Parallel()
	_, err := quiz.FromCSV("testdata/invalid.csv")
	if err == nil {
		t.Fatal("want error if CSV file is invalid.")
	}
}

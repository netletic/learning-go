package battery_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/battery"
)

func TestParsePmsetOutput_GetsChargePercent(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 23,
	}
	got, err := battery.ParsePmsetOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParsePmsetOutput_ReturnsErrorOnFailedParse(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset_invalid.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = battery.ParsePmsetOutput(string(data))
	if err == nil {
		t.Fatal("want error if pmset output fails to parse")
	}
}

func TestToJSON_GivesExpectedJSON(t *testing.T) {
	t.Parallel()
	batt := battery.Battery{
		Name:             "InternalBattery-0",
		ID:               35913827,
		ChargePercent:    100,
		TimeToFullCharge: "0:00",
		Present:          true,
	}
	wantBytes, err := os.ReadFile("testdata/battery.json")
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBytes)
	got := batt.ToJSON()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

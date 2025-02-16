package prom_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/prom"
)

func TestConfigFromYAML_CorrectlyParsesYAMLData(t *testing.T) {
	t.Parallel()
	want := prom.Config{
		Global: prom.GlobalConfig{
			ScrapeInterval:     time.Second * 15,
			EvaluationInterval: time.Second * 30,
			ScrapeTimeout:      time.Second * 10,
			ExternalLabels: map[string]string{
				"monitor": "codelab",
				"foo":     "bar",
			},
		},
	}
	got, err := prom.ConfigFromYAML("testdata/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

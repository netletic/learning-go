package weather_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/netletic/weather"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"weather": weather.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestParseResponse_CorrectlyParsesJSONData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 280.36,
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorGivenInvalidJSON(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather_invalid.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Fatal("want error parsing invalid response, got nil")
	}
}

func TestFormatURL_ReturnsCorrectURLForGivenInputs(t *testing.T) {
	t.Parallel()
	c := weather.NewClient("dummyAPIKey")
	location := "Dublin,IE"
	want := "https://api.openweathermap.org/data/2.5/weather?q=Dublin,IE&appid=dummyAPIKey"
	got := c.FormatURL(location)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestHTTPGet_SuccessfullyGetsFromLocalServer(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/weather.json")
		}))
	defer ts.Close()
	client := ts.Client()
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	want := http.StatusOK
	got := resp.StatusCode
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestGetWeather_ReturnsExpectedConditions(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/weather.json")
		}))
	defer ts.Close()
	c := weather.NewClient("dummyAPIKey")
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 280.36,
	}
	got, err := c.GetWeather("Ireland,IE")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCelsiusCorrectlyConversFahrenheitToCelsius(t *testing.T) {
	t.Parallel()
	input := weather.Temperature(274.15)
	want := 1.0
	got := input.Celsius()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

}

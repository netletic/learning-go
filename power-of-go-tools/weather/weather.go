package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const APIKeyEnvVar string = "OPENWEATHERMAP_API_KEY"

type Conditions struct {
	Summary     string
	Temperature Temperature
}
type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one weather element", data)
	}
	conditions := Conditions{
		Summary:     resp.Weather[0].Main,
		Temperature: Temperature(resp.Main.Temp),
	}
	return conditions, nil
}

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    "https://api.openweathermap.org",
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s",
		c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func Get(location, apiKey string) (Conditions, error) {
	c := NewClient(apiKey)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

const Usage = `Usage: weather LOCATION

Example: weather Ireland,IE
`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stdout, Usage)
		os.Exit(0)
	}
	apiKey := os.Getenv(APIKeyEnvVar)
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		os.Exit(1)
	}
	location := os.Args[1]
	conditions, err := Get(location, apiKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf("%s %.1f°C\n", conditions.Summary, conditions.Temperature.Celsius())
	return 0
}

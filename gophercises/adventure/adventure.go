package adventure

import (
	"encoding/json"
	"os"
)

type Adventure map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func FromJSON(path string) (Adventure, error) {
	f, err := os.Open(path)
	if err != nil {
		return Adventure{}, err
	}
	defer f.Close()

	var adventure Adventure
	err = json.NewDecoder(f).Decode(&adventure)
	if err != nil {
		return Adventure{}, err
	}

	return adventure, nil
}

package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type redirect struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func buildMap(redirects []redirect) map[string]string {
	redirectMap := make(map[string]string)
	for _, r := range redirects {
		redirectMap[r.Path] = r.URL
	}
	return redirectMap
}

func FromYAML(data []byte) (map[string]string, error) {
	redirects := []redirect{}
	err := yaml.Unmarshal(data, &redirects)
	if err != nil {
		return nil, err
	}
	return buildMap(redirects), nil

}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := FromYAML(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsToUrls, fallback), nil
}

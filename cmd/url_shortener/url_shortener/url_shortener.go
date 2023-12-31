package urlshortener

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type YamlRecord struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if link, ok := pathsToUrls[path]; ok != false {
			http.Redirect(w, r, link, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func parseYaml(yml []byte) ([]YamlRecord, error) {
	var parsedYaml []YamlRecord

	err := yaml.Unmarshal(yml, &parsedYaml)
	if err != nil {
		fmt.Println("err ", err)
		return nil, err
	}

	return parsedYaml, nil
}

func buildPathMap(data []YamlRecord) map[string]string {
	pathMap := make(map[string]string)

	for _, record := range data {
		pathMap[record.Path] = record.Url
	}

	return pathMap
}
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildPathMap(parsedYaml)

	return MapHandler(pathMap, fallback), nil
}

package handlers

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func YamlHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yaml)
	if err != nil {
		return nil, err
	}
	pathMap, _ := buildMap(parsedYaml.UrlMappings)
	return MapHandler(pathMap, fallback), nil
}

func parseYaml(yamlb []byte) (UrlMap, error) {
	urlMapping := UrlMap{}
	err := yaml.Unmarshal(yamlb, &urlMapping)
	if err != nil {
		return UrlMap{}, err
	}
	return urlMapping, nil
}

var internalMap = map[string]string{
	"/twitter":  "wwww.twitter",
	"/facebook": "www.facebook.com",
	"/eco":      "www.economist.com",
}

func buildMap(parsedYaml map[string]string) (map[string]string, error) {
	for k, v := range parsedYaml {
		internalMap[k] = v
	}
	return internalMap, nil
}

func MapHandler(pathMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Printf("Path: %s\n", path)
		if dest, ok := pathMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)

	}
}

type UrlMap struct {
	UrlMappings map[string]string `"yaml:UrlMappings"`
}

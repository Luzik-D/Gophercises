package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	urlshortener "github.com/Luzik-D/Gophercises/cmd/url_shortener/url_shortener"
)

func main() {
	yamlFilename := flag.String("yaml", "yml.yaml", "yaml file that containts info for redirections")

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/dota":           "https://www.dota2.com/home",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlFile, err := os.Open(*yamlFilename)
	if err != nil {
		panic(err)
	}
	defer yamlFile.Close()

	yaml, err := io.ReadAll(yamlFile)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshortener.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/chauhanr/go-building-blocks/url-shortner/handlers"
)

func main() {
	mux := defaultMux()
	yaml := `UrlMappings:
        "/gogl": "www.google.com"
        "/yah": "www.yahoo.com"
        "/hcl": "www.hcl.com"`

	yHandler, err := handlers.YamlHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":9090", yHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")

}

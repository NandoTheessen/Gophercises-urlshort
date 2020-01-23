package main

import (
	"fmt"
	"flag"
	"net/http"

	"github.com/NandoTheessen/Gophercises---urlshort/urlshort"
)

var file string
var format string

func init() {
	flag.StringVar(&file, "pathfile", "", "provide a json | yaml file composed of sequence of path & url mappings")
	flag.StringVar(&format, "format", "", "file format of the provided pathfile")
	flag.Parse()
}



func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	var yaml []byte
	var err error
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	if file != "" {
		yaml, err = reader.ProcessFile(file)
		
	} else {
		yamlstring := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
		yaml = []byte(yamlstring)
	}
	
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
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
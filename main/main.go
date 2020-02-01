package main

import (
	"fmt"
	"flag"
	"net/http"

	"github.com/NandoTheessen/Gophercises-urlshort/pkg/urlshort"
	"github.com/NandoTheessen/Gophercises-urlshort/pkg/reader"
)

var file string
var format string

func init() {
	flag.StringVar(&file, "file", "none", "provide a json | yaml file composed of sequence of path & url mappings")
	flag.StringVar(&format, "format", "yaml", "file format of the provided pathfile")
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

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	byteString, err := reader.ProcessFile(file)
	if err != nil {
		fmt.Println(err)
	}
	
	if format == "yaml" {
		yamlHandler, err := urlshort.YAMLHandler(byteString, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080 using yaml input")
		http.ListenAndServe(":8080", yamlHandler)
	} else if format == "json" {
		jsonHandler, err := urlshort.JSONHandler(byteString, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080 using json input")
		http.ListenAndServe(":8080", jsonHandler)
	} else {
		fmt.Println("The file format you've provided is invalid!")
		return
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
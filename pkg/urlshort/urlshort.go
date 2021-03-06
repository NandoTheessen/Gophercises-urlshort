package urlshort

import (
	"net/http"
	"fmt"
	"github.com/NandoTheessen/Gophercises-urlshort/pkg/reader"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func (rw http.ResponseWriter, req *http.Request) {
		
		path := req.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(rw, req, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw, req)
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
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(y []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathsToUrls, err := reader.ProcessYaml(y)

	if err != nil {
		fmt.Print(err)
	}
	return MapHandler(pathsToUrls, fallback), nil
}

//JSONHandler parses provided json & return http.Handlerfunc
func JSONHandler(j []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathsToUrls, err := reader.ProcessJSON(j)

	if err != nil {
		fmt.Print(err)
	}
	return MapHandler(pathsToUrls, fallback), nil
}
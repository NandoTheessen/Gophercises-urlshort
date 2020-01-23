package reader

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// ProcessFile takes a filename as argument and returns a bytearray with the 
// contents of that file or an error
func ProcessFile(filename string) (contents []byte, err error) {
	if filename == "" {
		yamlstring := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
		contents = []byte(yamlstring)
	} else {
		contents, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		
		}
	}
		return contents, nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}

// ProcessYaml takes in an array of bytes containing yaml file content & parses it
// into a pathsToUrls map
func ProcessYaml(yml []byte) (pathsToUrls map[string]string, err error) {
	var output []pathURL
	pathsToUrls = make(map[string]string)
	err = yaml.Unmarshal(yml, &output)
	if err != nil {
		return nil, err
	}

	for _, val := range output {
		pathsToUrls[val.Path] = val.URL
	}
	return pathsToUrls, nil
}

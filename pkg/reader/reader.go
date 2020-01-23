package reader

import (
	"io/ioutil"
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
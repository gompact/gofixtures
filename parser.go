package gofixtures

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func ParseDBConfFromYAML(filename string) (interface{}, error) {
	// check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = yaml.Unmarshal([]byte(input), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ParseYAML
func ParseYAML(filename string) (interface{}, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	var data interface{}

	err = yaml.Unmarshal([]byte(input), &data)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return data, nil
}

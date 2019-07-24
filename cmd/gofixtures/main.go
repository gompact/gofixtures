package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/schehata/gofixtures/entity"
	"github.com/schehata/gofixtures/logger"
	"github.com/schehata/gofixtures/v3"
)

var queries []string

const version = "3.0.0"

func main() {
	gf, err := gofixtures.New(entity.Config{})
	if err != nil {
		log.Fatal(err)
	}
	cmdArgs := os.Args
	switch cmdArgs[1] {
	case "version":
		log.Printf("version: %s\n", gf.Version())
		os.Exit(0)
	case "load":
		break
	case "clear":
		break
	default:
		log.Fatal("You must supply a command")
	}
	// read yaml config
	conf, err := ReadConfig(".gofixtures.yml")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// read input using CLI
	gf, err = gofixtures.New(conf)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if cmdArgs[1] == "load" {
		var workingPath = "./fixtures"
		if v := cmdArgs[2]; v != "" {
			workingPath = v
			if err != nil {
				log.Fatal(err)
			}
		}
		files, err := filesToParse(workingPath)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		err = gf.LoadFromFiles(files)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	} else if cmdArgs[1] == "clear" {
		gf.Clear()
	}

	logger.Success(fmt.Sprintf("Successfully inserted %d out of %d\n", 1, 1))
}

// FilesToParse checks if there is a filename is passed in the command line, If not,
// Check if a directory is passed, else, Expect to find a dir named "fixtures" to load
// files form.
// Returns a list of string of filenames
func filesToParse(givenPath string) ([]string, error) {
	files := []string{}
	// TOFIX: get current dir
	currentDir := ""
	p := path.Join(currentDir, givenPath)
	// TOFIX: check if this path is a file or a directory
	fileinfos, err := ioutil.ReadDir(p)
	if err != nil {
		return files, err
	}
	for _, f := range fileinfos {
		filename := path.Join(p, f.Name())
		files = append(files, filename)
	}
	return files, nil
}

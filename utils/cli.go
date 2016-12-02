package utils

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"fmt"
)

type CLI struct {
	Filename   string
	DBDriver   string
	DBConfFile string
	Directory  string
	DBConf     string
}

var defaultDirectory = "./fixtures"
var defaultDBConfigFile = "./db/dbconf.yml"

// ReadCommandLineFlags reads the options supplied in the comamnd line
func (c *CLI) ReadCommandLineFlags() {
	flag.StringVar(&c.Directory, "dir", "", "The path of the fixtures directory")
	flag.StringVar(&c.Filename, "file", "", "The path of a fixture file to load")
	flag.StringVar(&c.DBDriver, "driver", "postgres", "The database driver")
	flag.StringVar(&c.DBConfFile, "dbconffile", "", "The path of dbconf.yml file to load database configuration")
	flag.StringVar(&c.DBConf, "dbconf", "", "A string represents database configuration")

	flag.Parse()
}

// DatabaseConf generate the db configuration string. Make sure
// to call ReadCommandLineFlags() first, Because DatabaseConf()
// checks first if command line arguments has been passed
// if not it looks for dbconf.yml in ./db folder.
func (c CLI) DatabaseConf() (string, error) {
	if c.DBConf != "" {
		return c.DBConf, nil
	} else if c.DBConfFile != "" {
		data, err := ParseDBConfFromYAML(c.DBConfFile)
		if err != nil {
			return "", err
		}
		result := data.(map[interface{}]interface{})
		c.DBDriver = result["driver"].(string)
		c.DBConf = result["open"].(string)
	} else {
		// look for ./db/dbconf.yml
		log.Print("Expecting to find dbconf.yml in db/")
		c.DBConfFile = "db/dbconf.yml"
		data, err := ParseDBConfFromYAML(c.DBConfFile)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		result := data.(map[interface{}]interface{})
		c.DBDriver = result["driver"].(string)
		c.DBConf = result["open"].(string)
	}
	// check if c.DBConf has value, otherwise we failed to find any configuration
	if c.DBConf == "" {
		errStr := `Failed to find any configurations for database. Please pass db conf through commandline or yaml file.`
		return "", errors.New(errStr)
	}
	return c.DBConf, nil
}

// FilesToParse checks if there is a filename is passed in the command line, If not,
// Check if a directory is passed, Else, Expect to find a dir named "fixtures" to load
// files form.
// Returns a list of string of filenames
func (c *CLI) FilesToParse() []string {
	files := []string{}
	if c.Filename != "" {
		filename := fmt.Sprintf("%s/%s", c.Directory, c.Filename)
		files = append(files, filename)
	} else if c.Directory != "" {
		fileinfos, err := ioutil.ReadDir(c.Directory)
		if err != nil {
			panic(err)
		}
		for _, f := range fileinfos {
			filename := fmt.Sprintf("%s/%s", c.Directory, f.Name())
			files = append(files, filename)
		}
	} else {
		fileinfos, err := ioutil.ReadDir("./fixtures")
		if err != nil {
			panic(err)
		}
		c.Directory = "fixtures"
		for _, f := range fileinfos {
			filename := fmt.Sprintf("%s/%s", c.Directory, f.Name())
			files = append(files, filename)
		}
	}
	return files
}

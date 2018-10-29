package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ishehata/gofixtures/entity"
)

// CLI implements Feeder interface
type CLI struct {
	filename   string
	configFile string
	directory  string
	currentDir string
}

func init() {
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// New creates a new instace of the CLI feeder
func New() *CLI {
	feeder := &CLI{}
	feeder.readCommandLineFlags()
	feeder.currentDir, _ = filepath.Abs("./")
	return feeder
}

const defaultFixturesDirectory = "fixtures"
const defaultConfigFile = ".gofixtures.yaml"

// readCommandLineFlags reads the options supplied in the comamnd line
func (cli *CLI) readCommandLineFlags() {
	flag.StringVar(&cli.directory, "dir", defaultFixturesDirectory, "The path of the fixtures directory")
	flag.StringVar(&cli.filename, "file", "", "The path of a fixture file to load")
	flag.StringVar(&cli.configFile, "config", defaultConfigFile, "The path of config file to load configurations from")

	flag.Parse()
}

// ReadConfig reades the configurations file. Make sure
// to call ReadCommandLineFlags() first.
func (cli *CLI) ReadConfig() (entity.ConfigInput, error) {
	// determine the file type, e.g: yaml or json
	ext := filepath.Ext(cli.configFile)

	input := entity.ConfigInput{
		Type: ext,
	}

	cli.Info("reading configuration file...")
	f, err := os.Open(cli.configFile)
	if err != nil {
		return input, err
	}
	cli.Success("configuration file has been loaded successfully")

	input.Data = f

	return input, nil
}

// extractFilename returns a filename that could be used as a tablename
// in case of csv files. it splits the paths and removes the file extension
func extractFilename(filePath string) string {
	splitted := strings.Split(filePath, "/")
	filenameWithExt := splitted[len(splitted)-1]
	fileNameSplitted := strings.Split(filenameWithExt, ".")
	filename := fileNameSplitted[0]
	return filename
}

// GetInput reads the files selected by the user
// and return them as inputs
func (cli *CLI) GetInput() ([]entity.Input, error) {
	var inputs []entity.Input
	// get list of files that will be parsed
	files, err := cli.filesToParse()
	if err != nil {
		return inputs, err
	}
	inputs = make([]entity.Input, len(files))
	cli.Print(fmt.Sprintf("found %d fixture files", len(files)))
	for i, file := range files {
		cli.Print(fmt.Sprintf("reading fixture: %s", file))
		f, err := os.Open(file)
		if err != nil {
			return inputs, err
		}
		ext := filepath.Ext(file)
		input := entity.Input{
			Filename: extractFilename(file),
			Type:     ext,
			Data:     f,
		}
		inputs[i] = input
	}
	return inputs, nil
}

// FilesToParse checks if there is a filename is passed in the command line, If not,
// Check if a directory is passed, Else, Expect to find a dir named "fixtures" to load
// files form.
// Returns a list of string of filenames
func (cli *CLI) filesToParse() ([]string, error) {
	files := []string{}
	if cli.filename != "" {
		filename := path.Join(cli.currentDir, cli.filename)
		files = append(files, filename)
	} else {
		// if no file is selected then read all files in the fixtures directory
		// if the user didn't specify a directly, the default one will be used
		// which was set in flag parsing section
		fileinfos, err := ioutil.ReadDir(cli.directory)
		if err != nil {
			return files, err
		}
		for _, f := range fileinfos {
			filename := path.Join(cli.directory, f.Name())
			files = append(files, filename)
		}
	}
	return files, nil
}

// Print logs text to the end user
func (cli *CLI) Print(text string) {
	info(text)
}

// Info logs text to the end user
func (cli *CLI) Info(text string) {
	info(text)
}

// Debug logs text to the end user
func (cli *CLI) Debug(text string) {
	debug(text)
}

// Warning logs text to the end user
func (cli *CLI) Warning(text string) {
	warn(text)
}

// Success logs text to the end user
func (cli *CLI) Success(text string) {
	success(text)
}

// Error prints and error to the user, exists if its a fatal error
func (cli *CLI) Error(err error, fatal bool) {
	// txt := fmt.Sprintf("%#v", err)
	errorF(err.Error())
	if fatal {
		os.Exit(1)
	}
}

package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/schehata/gofixtures/v3/entity"
	"github.com/schehata/gofixtures/v3/logger"
)

// Feeder is a file reader that implements Feeder interface
type Feeder struct {
	Config Config
}

// Config holds the configuration information for the file feeder
type Config struct {
	CurrentDir string
}

// New returns a new instance of the file feeder
func New() *Feeder {
	return &Feeder{}
}

// Read converts data files into Fixture consumable data
func (feeder *Feeder) Read(files []string) ([]entity.Input, error) {
	inputs := make([]entity.Input, len(files))
	for i, file := range files {
		logger.Info(fmt.Sprintf("reading fixture: %s", file))
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

// extractFilename returns a filename that could be used as a tablename
// in case of csv files. it splits the paths and removes the file extension
func extractFilename(filePath string) string {
	splitted := strings.Split(filePath, "/")
	filenameWithExt := splitted[len(splitted)-1]
	fileNameSplitted := strings.Split(filenameWithExt, ".")
	filename := fileNameSplitted[0]
	return filename
}

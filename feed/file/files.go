package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ishehata/gofixtures/entity"
	"github.com/ishehata/gofixtures/logger"
)

// Feeder is a file reader that implements Feeder interface
type Feeder struct {
	Config Config
}

// Config holds the configuration information for the file feeder
type Config struct {
	Files      []string
	CurrentDir string
}

// New returns a new instance of the file feeder
func New() *Feeder {
	return &Feeder{}
}

func (feeder *Feeder) Read() ([]entity.Input, error) {
	inputs := make([]entity.Input, len(feeder.Config.Files))
	for i, file := range feeder.Config.Files {
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

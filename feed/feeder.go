package feed

import (
	"github.com/schehata/gofixtures/v3/entity"
)

// Feeder interface defines methods needs to be implemented
// by different interfaces that could act as the input for gofixtures
type Feeder interface {
	// Read retrieves the list of fixtures that should be loaded
	// into the datastore
	Read(files []string) ([]entity.Input, error)
}

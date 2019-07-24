package feed

import (
	"github.com/schehata/gofixtures/v3/entity"
)

// Feeder interface defines methods needs to be implemented
// by different interfaces that could act as the input for gofixtures
type Feeder interface {
	// GetInput retrieves the list of fixtures that should be loaded
	// into the datastore
	GetInput() ([]entity.Input, error)
	// Print sends/prints something to the end user
	Info(string)
	Warning(string)
	Debug(string)
	// Error prints an error to the end user and it could posibbly
	// end the session if its a fatal error
	Error(error, bool)
}

package dal

import "github.com/ishehata/gofixtures/v3/entity"

// Datastore interface defines the methods needed to be implemented
// by different database drivers
type Datastore interface {
	Connect() error
	Insert(fixture entity.Fixture) error
	Close()
}

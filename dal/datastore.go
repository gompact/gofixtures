package dal

import "github.com/ishehata/gofixtures/entity"

// Datastore interface defines the methods needed to be implemented
// by different database drivers
type Datastore interface {
	Connect(config entity.DBConfig) error
	Insert(entity.Fixture) error
	Close()
}

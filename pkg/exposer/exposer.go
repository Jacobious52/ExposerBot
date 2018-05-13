package exposer

import "github.com/Jacobious52/expose/pkg/storage"

// Exposer is an interface for something that can expose metrics from a datastore
type Exposer interface {
	// Setup will be run once apon initialisation
	// Here is where auxiliary data sources can be loaded
	Setup() error
	// Expose will run every call to the registered plugin
	// Expose should return a formatted string of the output
	Expose(storage.ReadStore) (string, error)
}

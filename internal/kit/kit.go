package kit

import (
	"context"

	"github.com/thefarmhub/farmhub-cli/internal/model"
)

type Kit interface {
	// SetPort specifies where it should be operating when flashing and monitoring
	SetPort(port string)

	// SetPath specifies the path to the sketch to be flashed
	SetPath(path string) error

	// Init initializes the kit, installing libraries and other dependencies
	Init() error

	// Upload flashes the sketch to the specified port
	Upload() error

	// GenerateCode generates the code necessary for this kit
	GenerateCode(sensor *model.Sensor) (string, error)

	// Monitor starts monitoring the specified port
	Monitor(ctx context.Context) error
}

// KitConstructor is a function that returns a new Kit and allows
// them to be registered in the map below without initializing
// them immediately
type KitConstructor func() Kit

// availableKits is a map of kit names to their constructors
var availableKits = make(map[string]KitConstructor)

// GetKitNames returns a list of available kit names to be used
// in lists and interactive menus
func GetKitNames() []string {
	var names []string

	for name := range availableKits {
		names = append(names, name)
	}

	return names
}

// GetKitByName returns a new kit instance by its name
func GetKitByName(name string) Kit {
	return availableKits[name]()
}

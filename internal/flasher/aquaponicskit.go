package flasher

import (
	"github.com/arduino/arduino-cli/configuration"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/thefarmhub/farmhub-cli/internal/arduino"
)

type AquaponicsKit struct {
	arduino *arduino.Arduino
	path    string
	port    string
}

func NewAquaponicsKit() *AquaponicsKit {
	configuration.Settings = configuration.Init("")

	return &AquaponicsKit{
		arduino: arduino.NewArduino("esp32:esp32:featheresp32"),
	}
}

func (e *AquaponicsKit) SetPort(port string) {
	e.port = port
}

func (e *AquaponicsKit) Init() error {
	configuration.Settings.Set("board_manager.additional_urls.0", "https://raw.githubusercontent.com/espressif/arduino-esp32/gh-pages/package_esp32_index.json")

	_, err := e.arduino.PlatformInstall(&rpc.PlatformInstallRequest{
		Architecture: "esp32",
		PlatformPackage: "esp32",
	})

	err = e.arduino.InstallLibrary(&rpc.LibraryInstallRequest{
		Name: "PubSubClient",
		Version: "2.8.0",
	})
	if err != nil {
		return err
	}

	err = e.arduino.GitLibraryInstall(&rpc.GitLibraryInstallRequest{
		Url: "https://github.com/Atlas-Scientific/Ezo_I2c_lib.git#dbb83f3",
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *AquaponicsKit) Upload() error {
	err := e.arduino.Compile(e.path)
	if err != nil {
		return err
	}

	return e.arduino.Upload(e.port, e.path)
}

func (e *AquaponicsKit) SetPath(path string) {
	e.path = path
}

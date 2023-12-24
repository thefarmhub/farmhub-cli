package flasher

import (
	"github.com/arduino/arduino-cli/configuration"
	"github.com/thefarmhub/farmhub-cli/internal/arduino"
)

type esp32 struct {
	arduino *arduino.Arduino
	path	string
	port	string
}

func NewESP32(fbqn string) Flasher {
	configuration.Settings = configuration.Init("")

	return &esp32{
		arduino: arduino.NewArduino(fbqn),
	}
}

func (e *esp32) SetPort(port string) Flasher {
	e.port = port
	return e
}

func (e *esp32) Init() error {
	configuration.Settings.Set("board_manager.additional_urls.0", "https://raw.githubusercontent.com/espressif/arduino-esp32/gh-pages/package_esp32_index.json")

	return nil
}

func (e *esp32) Upload() error {
	err := e.arduino.Compile(e.path)
	if err != nil {
		return err
	}

	err = e.arduino.Upload(e.port, e.path)

	return err
}

func (e *esp32) SetPath(path string) Flasher {
	e.path = path
	return e
}

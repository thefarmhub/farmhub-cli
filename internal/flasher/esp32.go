package flasher

import (
	"github.com/arduino/arduino-cli/configuration"
	"github.com/thefarmhub/farmhub-cli/internal/arduino"
)

type ESP32 struct {
	arduino *arduino.Arduino
	path	string
	port	string
}

func NewESP32(fbqn string) *ESP32 {
	configuration.Settings = configuration.Init("")

	return &ESP32{
		arduino: arduino.NewArduino(fbqn),
	}
}

func (e *ESP32) SetPort(port string) {
	e.port = port
}

func (e *ESP32) Init() error {
	configuration.Settings.Set("board_manager.additional_urls.0", "https://raw.githubusercontent.com/espressif/arduino-esp32/gh-pages/package_esp32_index.json")

	return nil
}

func (e *ESP32) Upload() error {
	err := e.arduino.Compile(e.path)
	if err != nil {
		return err
	}

	return e.arduino.Upload(e.port, e.path)
}

func (e *ESP32) SetPath(path string) {
	e.path = path
}

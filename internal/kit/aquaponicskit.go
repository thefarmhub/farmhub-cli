package kit

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/arduino/arduino-cli/commands/monitor"
	"github.com/arduino/arduino-cli/configuration"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/thefarmhub/farmhub-cli/internal/arduino"
	"github.com/thefarmhub/farmhub-cli/internal/arduino/cli/feedback"
	"go.bug.st/cleanup"
)

type AquaponicsKit struct {
	arduino *arduino.Arduino
	path    string
	port    string
	fqbn   string
}

func NewAquaponicsKit() Kit {
	configuration.Settings = configuration.Init("")

	fqbn := "esp32:esp32:featheresp32"

	a := arduino.NewArduino()
	a.SetFQBN(fqbn)

	return &AquaponicsKit{
		arduino: a,
		fqbn: fqbn,
	}
}

func (e *AquaponicsKit) SetPort(port string) {
	e.port = port
}

func (e *AquaponicsKit) Init() error {
	configuration.Settings.Set("board_manager.additional_urls.0", "https://raw.githubusercontent.com/espressif/arduino-esp32/gh-pages/package_esp32_index.json")

	_, err := e.arduino.PlatformInstall(&rpc.PlatformInstallRequest{
		Architecture:    "esp32",
		PlatformPackage: "esp32",
	})

	err = e.arduino.InstallLibrary(&rpc.LibraryInstallRequest{
		Name:    "PubSubClient",
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

func (e *AquaponicsKit) SetPath(path string) error {
	newPath, err := arduino.PrepareSketch(path)
	if err != nil {
		return err
	}

	e.path = newPath

	return nil
}

func (e *AquaponicsKit) Monitor(ctx context.Context) error {
	feedback.SetFormat(feedback.Text)

	configuration := &rpc.MonitorPortConfiguration{}
	portProxy, _, err := monitor.Monitor(context.Background(), &rpc.MonitorRequest{
		Instance:          e.arduino.GetInstance(),
		Port:              &rpc.Port{Address: e.port, Protocol: "serial"},
		Fqbn:              e.fqbn,
		PortConfiguration: configuration,
	})
	if err != nil {
		feedback.FatalError(err, feedback.ErrGeneric)
	}
	defer portProxy.Close()

	feedback.Print(fmt.Sprintf("Connected to %s! Press CTRL-C to exit.", e.port))

	ttyIn, ttyOut, err := feedback.InteractiveStreams()
	if err != nil {
		feedback.FatalError(err, feedback.ErrGeneric)
	}

	ctx, cancel := cleanup.InterruptableContext(context.Background())

	go func() {
		_, err := io.Copy(ttyOut, portProxy)
		if err != nil && !errors.Is(err, io.EOF) {
			feedback.Print(fmt.Sprintf("Port closed: %v", err))
		}
		cancel()
	}()
	go func() {
		_, err := io.Copy(portProxy, ttyIn)
		if err != nil && !errors.Is(err, io.EOF) {
			feedback.Print(fmt.Sprintf("Port closed: %v", err))
		}
		cancel()
	}()

	<-ctx.Done()
	return nil
}

func init() {
	availableKits["aquaponics-kit-v1"] = NewAquaponicsKit
}

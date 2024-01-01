package kit

import (
	"bytes"
	"context"
	"strings"
	"text/template"

	_ "embed"

	"github.com/arduino/arduino-cli/configuration"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/thefarmhub/farmhub-cli/internal/arduino"
	"github.com/thefarmhub/farmhub-cli/internal/datacompletion"
	"github.com/thefarmhub/farmhub-cli/internal/model"
)

//go:embed templates/hydroponicskitv2.ino
var hydroponicsKitV2Template string

type HydroponicsKitV2 struct {
	arduino *arduino.Arduino
	path    string
	port    string
}

func NewHydroponicsKitV2() Kit {
	configuration.Settings = configuration.Init("")

	fqbn := "esp32:esp32:featheresp32"

	a := arduino.NewArduino()
	a.SetFQBN(fqbn)

	return &HydroponicsKitV2{
		arduino: a,
	}
}

func (e *HydroponicsKitV2) SetPort(port string) {
	e.port = port
}

func (e *HydroponicsKitV2) Init() error {
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

func (e *HydroponicsKitV2) Upload() error {
	err := e.arduino.Compile(e.path)
	if err != nil {
		return err
	}

	return e.arduino.Upload(e.port, e.path)
}

func (e *HydroponicsKitV2) SetPath(path string) error {
	newPath, err := arduino.PrepareSketch(path)
	if err != nil {
		return err
	}

	e.path = newPath

	return nil
}

func (e *HydroponicsKitV2) Monitor(ctx context.Context) error {
	return e.arduino.Monitor(ctx, e.port)
}

func (e *HydroponicsKitV2) GenerateCode(sensor *model.Sensor) (string, error) {
	tmpl, err := template.New("code").Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
	}).Parse(hydroponicsKitV2Template)
	if err != nil {
		return "", err
	}

	type ConfigVariables struct {
		WiFiSSID                 string `datatype:"ssid"`
		WiFiPassword             string `datatype:"password"`
		TopicPH                  string `datatype:"topic" metric:"PH" name:"pH"`
		TopicEC                  string `datatype:"topic" metric:"ELECTRICAL_CONDUCTIVITY" name:"Electrical Conductivity"`
		TopicTEMP                string `datatype:"topic" metric:"WATER_TEMPERATURE" name:"Water Temperature"`
		ThingName                string
		CertificatePEM           string
		CertificatePrivateKey    string
		RootCertificateAuthority string
		IotEndpoint              string
	}

	vars := ConfigVariables{
		IotEndpoint:              IotEndpoint,
		CertificatePEM:           sensor.IoTCertificatePem,
		CertificatePrivateKey:    sensor.IoTCertificatePrivateKey,
		RootCertificateAuthority: sensor.IoTRootCertificateAuthority,
		ThingName:                sensor.IoTThingName,
	}

	datacompletion.Complete(&vars, sensor)

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, vars)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func init() {
	availableKits["hydroponics-kit-v2"] = NewHydroponicsKitV2
}

package kit

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	_ "embed"

	"github.com/arduino/arduino-cli/configuration"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/thefarmhub/farmhub-cli/internal/arduino"
	"github.com/thefarmhub/farmhub-cli/internal/model"
)

//go:embed templates/aquaponicskitv1.ino
var aquaponicsKitV1Template string

type AquaponicsKitV1 struct {
	arduino *arduino.Arduino
	path    string
	port    string
}

func NewAquaponicsKitV1() Kit {
	configuration.Settings = configuration.Init("")

	fqbn := "esp32:esp32:featheresp32"

	a := arduino.NewArduino()
	a.SetFQBN(fqbn)

	return &AquaponicsKitV1{
		arduino: a,
	}
}

func (e *AquaponicsKitV1) SetPort(port string) {
	e.port = port
}

func (e *AquaponicsKitV1) Init() error {
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

func (e *AquaponicsKitV1) Upload() error {
	err := e.arduino.Compile(e.path)
	if err != nil {
		return err
	}

	return e.arduino.Upload(e.port, e.path)
}

func (e *AquaponicsKitV1) SetPath(path string) error {
	newPath, err := arduino.PrepareSketch(path)
	if err != nil {
		return err
	}

	e.path = newPath

	return nil
}

func (e *AquaponicsKitV1) GenerateCode(sensor *model.Sensor) (string, error) {
	fmt.Println(aquaponicsKitV1Template)
	tmpl, err := template.New("code").Parse(aquaponicsKitV1Template)
	if err != nil {
		return "", err
	}

	type ConfigVariables struct {
		WiFiSSID                 string
		WiFiPassword             string
		TopicPH                  string
		TopicEC                  string
		TopicDO                  string
		TopicTEMP                string
		TopicHUM                 string
		TopicCO2                 string
		ThingName                string
		CertificatePEM           string
		CertificatePrivateKey    string
		RootCertificateAuthority string
		IotEndpoint              string
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, ConfigVariables{
		IotEndpoint:              "iot.farmhub.ag",
		CertificatePEM:           sensor.IoTCertificatePem,
		CertificatePrivateKey:    sensor.IoTCertificatePrivateKey,
		RootCertificateAuthority: sensor.IoTRootCertificateAuthority,
		ThingName:                sensor.IoTThingName,
	})
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func (e *AquaponicsKitV1) Monitor(ctx context.Context) error {
	return e.arduino.Monitor(ctx, e.port)
}

func init() {
	availableKits["aquaponics-kit-v1"] = NewAquaponicsKitV1
}

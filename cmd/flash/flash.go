package flash

import (
	"bytes"
	"context"
	"os"

	"github.com/arduino/arduino-cli/commands/compile"
	"github.com/arduino/arduino-cli/commands/lib"
	"github.com/arduino/arduino-cli/commands/upload"
	"github.com/arduino/arduino-cli/configuration"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/internal/cli/feedback"
	"github.com/thefarmhub/farmhub-cli/internal/cli/instance"
	"go.bug.st/serial"
)

const FBQN = "esp32:esp32:featheresp32"

func NewFlashCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "flash",
		Short: "Flashes to ESP32",
		RunE: func(cmd *cobra.Command, args []string) error {
			ports, err := serial.GetPortsList()
			if err != nil {
				pterm.Fatal.Println("Error fetching port list:", err)
			}

			if len(ports) == 0 {
				pterm.Warning.Println("No serial ports found! Please connect your device and try again.")
				return nil
			}

			selectedPort, err := pterm.DefaultInteractiveSelect.WithOptions(ports).Show()
			if err != nil {
				pterm.Fatal.Println("Error selecting port:", err)
			}

			pterm.Success.Println("You selected:", selectedPort)

			spinnerInit, _ := pterm.DefaultSpinner.Start("Initializing configuration...")
			configuration.Settings = configuration.Init(configuration.FindConfigFileInArgs(os.Args))

			configuration.Settings.Set("board_manager.additional_urls.0", "https://raw.githubusercontent.com/espressif/arduino-esp32/gh-pages/package_esp32_index.json")

			i := instance.CreateAndInit()
			spinnerInit.Success("Configuration initialized")

			// Install required libraries
			_ = installLibrary(i, "ArduinoJson", "6.21.2")

			spinnerCompile, _ := pterm.DefaultSpinner.Start("Compiling sketch...")
			ctx := context.Background()
			compileReq := &rpc.CompileRequest{
				Fqbn:       FBQN,
				Instance:   i,
				SketchPath: "/Users/albanda/Projects/farmhub/hardware-starter-kits/scientific-atlas/v2/aquaponics-kit",
			}

			var outStream, errStream bytes.Buffer
			_, err = compile.Compile(ctx, compileReq, &outStream, &errStream, nil)
			if err != nil {
				spinnerCompile.Fail(err.Error())
				return nil
			}

			spinnerCompile.Success("Successfully compiled sketch")

			spinnerUpload, _ := pterm.DefaultSpinner.Start("Uploading sketch...")
			uploadReq := &rpc.UploadRequest{
				Fqbn:       FBQN,
				Port:       &rpc.Port{Address: selectedPort},
				Instance:   i,
				SketchPath: compileReq.SketchPath,
			}

			_, uploadErr := upload.Upload(ctx, uploadReq, &outStream, &errStream)
			if uploadErr != nil {
				spinnerUpload.Fail(uploadErr.Error())
				return uploadErr
			}

			spinnerUpload.Success("Successfully uploaded sketch")

			return nil
		},
	}

	return cmd
}

func installLibrary(i *rpc.Instance, libraryName, libraryVersion string) error {
	req := &rpc.LibraryInstallRequest{
		Instance: i,
		Name:     libraryName,
		Version:  libraryVersion,
	}

	return lib.LibraryInstall(context.TODO(), req, feedback.ProgressBar(), feedback.TaskProgress())
}

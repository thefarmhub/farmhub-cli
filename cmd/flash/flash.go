package flash

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/internal/kit"
	"go.bug.st/serial"
)

func NewFlashCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "flash",
		Short: "Flashes to hardware",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			selectedPort := mustSelectPort()
			hardware := mustSelectKit()
			hardware.SetPort(selectedPort)

			err := hardware.SetPath(args[0])
			if err != nil {
				return err
			}

			spinnerInit, _ := pterm.DefaultSpinner.Start("Setting up configuration...")
			err = hardware.Init()
			if err != nil {
				spinnerInit.Fail(err.Error())
				return err
			}

			spinnerInit.Success("Configuration initialized")

			spinnerUpload, _ := pterm.DefaultSpinner.Start("Flashing...")
			err = hardware.Upload()
			if err != nil {
				spinnerUpload.Fail(err.Error())
				return err
			}

			spinnerUpload.Success("Successfully uploaded sketch")

			return nil
		},
	}

	return cmd
}

func mustSelectPort() string {
	ports, err := serial.GetPortsList()
	if err != nil {
		pterm.Fatal.Println("Could not list ports", err)
	}

	if len(ports) == 0 {
		pterm.Fatal.Println("No serial ports found! Please connect your device and try again.")
	}

	selectedPort, err := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select the port for sensor: ").
		WithOptions(ports).
		Show()
	if err != nil {
		pterm.Fatal.Println("Error selecting port:", err)
	}

	pterm.Success.Println("Selected", selectedPort)

	return selectedPort
}

func mustSelectKit() kit.Kit {
	selectedKit, err := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select the sensor: ").
		WithOptions(kit.GetKitNames()).
		Show()
	if err != nil {
		pterm.Fatal.Println("Error selecting kit:", err)
	}

	return kit.GetKitByName(selectedKit)
}

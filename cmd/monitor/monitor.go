package monitor

import (
	"context"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/internal/kit"
	"go.bug.st/serial"
)

func NewMonitorCommand() *cobra.Command {
	var monitorCmd = &cobra.Command{
		Use:   "monitor",
		Short: "Monitor serial port data",
		Run: func(cmd *cobra.Command, args []string) {
			port := mustSelectPort()
			hardware := mustSelectKit()
			hardware.SetPort(port)

			ctx := context.Background()
			hardware.Monitor(ctx)
		},
	}

	return monitorCmd
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

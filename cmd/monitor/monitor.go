package monitor

import (
	"context"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/internal/kit"
	"go.bug.st/serial"
)

var (
	kitFlag string
	port string
)

func NewMonitorCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "monitor",
		Short: "Monitors the serial port",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			selectedPort := mustSelectPort()
			hardware := mustSelectKit()
			hardware.SetPort(selectedPort)

			ctx := context.Background()
			return hardware.Monitor(ctx)
		},
	}

	cmd.Flags().StringVarP(&kitFlag, "kit", "k", "", "Select the kit to use")
	cmd.Flags().StringVarP(&port, "port", "p", "", "Select the port to use")

	return cmd
}

func mustSelectPort() string {
	if port != "" {
		return port
	}

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
	var err error
	selectedKit := kitFlag

	if selectedKit == "" {
		selectedKit, err = pterm.DefaultInteractiveSelect.
			WithDefaultText("Select the sensor: ").
			WithOptions(kit.GetKitNames()).
			Show()
		if err != nil {
			pterm.Fatal.Println("Error selecting kit:", err)
		}
	}

	return kit.GetKitByName(selectedKit)
}

package core

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/cmd/flash"
	"github.com/thefarmhub/farmhub-cli/cmd/monitor"
	versioncmd "github.com/thefarmhub/farmhub-cli/cmd/version"
	"github.com/thefarmhub/farmhub-cli/internal/version"
)

var verbose bool

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "farmhub",
		Version: version.Version,
		Long: `
 _____                    _   _       _
|  ___|_ _ _ __ _ __ ___ | | | |_   _| |__
| |_ / _  | '__| '_ ' _ \| |_| | | | | '_ \
|  _| (_| | |  | | | | | |  _  | |_| | |_) |
|_|  \__,_|_|  |_| |_| |_|_| |_|\__,_|_.__/

FarmHub CLI is a utility program created by FarmHub that helps you
configure and publish your IoT sensor data to your dashboard.
`,
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(flash.NewFlashCommand())
	rootCmd.AddCommand(monitor.NewMonitorCommand())
	rootCmd.AddCommand(versioncmd.NewVersionCommand())

	return rootCmd
}

func init() {
	cobra.OnInitialize(func() {
		if verbose {
			log.SetLevel(log.InfoLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}
	})
}

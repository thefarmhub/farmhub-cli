package core

import (
	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/cmd/flash"
	versioncmd "github.com/thefarmhub/farmhub-cli/cmd/version"
	"github.com/thefarmhub/farmhub-cli/internal/version"
)

var cfgDir string
var verbose bool

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "farmhub",
		Short:   "The loyal and hard working data dog that relays IoT sensor data to the cloud",
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

	rootCmd.AddCommand(flash.NewFlashCommand())
	rootCmd.AddCommand(versioncmd.NewVersionCommand())

	return rootCmd
}

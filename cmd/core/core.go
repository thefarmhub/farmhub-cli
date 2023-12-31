package core

import (
	"path"

	log "github.com/sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thefarmhub/farmhub-cli/cmd/flash"
	"github.com/thefarmhub/farmhub-cli/cmd/login"
	"github.com/thefarmhub/farmhub-cli/cmd/logout"
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
	rootCmd.AddCommand(login.NewLoginCommand())
	rootCmd.AddCommand(logout.NewLogoutCommand())
	rootCmd.AddCommand(versioncmd.NewVersionCommand())
	rootCmd.AddCommand(monitor.NewMonitorCommand())

	return rootCmd
}

func init() {
	cobra.OnInitialize(func() {
		if verbose {
			log.SetLevel(log.InfoLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}

		home, _ := homedir.Dir()
		viper.SetConfigFile(path.Join(home, ".farmhub.yaml"))
		viper.SetConfigType("yaml")
		viper.ReadInConfig()
	})
}

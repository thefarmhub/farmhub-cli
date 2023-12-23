package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thefarmhub/farmhub-cli/internal/version"
)

func NewVersionCommand() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of FarmHub",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("FarmHub CLI version: %s\n", version.Version)
			version.CheckLatestVersion()
		},
	}

	return versionCmd
}

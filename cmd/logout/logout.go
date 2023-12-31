package logout

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewLogoutCommand() *cobra.Command {
	var logoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "Destroy the authenticated session.",
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set("auth.token", "")
			viper.WriteConfig()

			fmt.Println("You have been successfully logged out.")
		},
	}

	return logoutCmd
}

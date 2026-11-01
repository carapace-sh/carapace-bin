package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of an npm registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	logoutCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	logoutCmd.Flags().String("registry", "", "The registry to log out of")
	rootCmd.AddCommand(logoutCmd)
}

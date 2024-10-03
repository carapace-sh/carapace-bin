package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove an API token from the registry locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	logoutCmd.Flags().BoolP("help", "h", false, "Print help")
	logoutCmd.Flags().String("registry", "", "Registry to use")
	rootCmd.AddCommand(logoutCmd)
}

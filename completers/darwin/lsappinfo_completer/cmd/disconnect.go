package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from launchservicesd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disconnectCmd).Standalone()
	rootCmd.AddCommand(disconnectCmd)
}

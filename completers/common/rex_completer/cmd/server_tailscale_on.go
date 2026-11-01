package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Serve the API over the tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleOnCmd).Standalone()

	serverTailscaleCmd.AddCommand(serverTailscaleOnCmd)
}

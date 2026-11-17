package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscaleOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Stop serving the API over the tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscaleOffCmd).Standalone()

	serverTailscaleOffCmd.Flags().Bool("json", false, "output the resulting state as JSON")

	serverTailscaleCmd.AddCommand(serverTailscaleOffCmd)
}

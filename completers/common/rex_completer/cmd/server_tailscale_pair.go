package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverTailscalePairCmd = &cobra.Command{
	Use:   "pair",
	Short: "Show a QR code to pair another device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverTailscalePairCmd).Standalone()

	serverTailscalePairCmd.Flags().Bool("json", false, "output the address as JSON")

	serverTailscaleCmd.AddCommand(serverTailscalePairCmd)
}

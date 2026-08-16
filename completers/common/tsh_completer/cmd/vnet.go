package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vnetCmd = &cobra.Command{
	Use:   "vnet",
	Short: "Start Teleport VNet, a virtual network for TCP application access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vnetCmd).Standalone()

	vnetCmd.Flags().Bool("diag", false, "Run diagnostics after starting VNet.")
	vnetCmd.Flags().Bool("no-diag", false, "Run diagnostics after starting VNet.")
	vnetCmd.Flag("diag").Hidden = true
	vnetCmd.Flag("no-diag").Hidden = true
	rootCmd.AddCommand(vnetCmd)
}

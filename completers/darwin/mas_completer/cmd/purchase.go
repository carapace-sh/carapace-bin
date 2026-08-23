package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var purchaseCmd = &cobra.Command{
	Use:     "purchase",
	Aliases: []string{"get"},
	Short:   "Get and install free apps from the App Store",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purchaseCmd).Standalone()
	purchaseCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	purchaseCmd.Flags().Bool("force", false, "Force reinstall")
	rootCmd.AddCommand(purchaseCmd)
}

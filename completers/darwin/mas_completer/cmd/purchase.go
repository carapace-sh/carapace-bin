package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var purchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase and install an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purchaseCmd).Standalone()
	rootCmd.AddCommand(purchaseCmd)
}

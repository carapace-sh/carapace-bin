package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Purchase Vercel products for your team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buyCmd).Standalone()

	rootCmd.AddCommand(buyCmd)
}

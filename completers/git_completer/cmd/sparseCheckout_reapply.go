package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_reapplyCmd = &cobra.Command{
	Use:   "reapply",
	Short: "Reapply the existing sparse directory specifications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_reapplyCmd).Standalone()

	sparseCheckoutCmd.AddCommand(sparseCheckout_reapplyCmd)
}

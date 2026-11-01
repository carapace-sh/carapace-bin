package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the sparse-checkout file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_initCmd).Standalone()

	sparseCheckout_initCmd.Flags().Bool("no-cone", false, "use non-cone mode")
	sparseCheckout_initCmd.Flags().Bool("no-sparse-index", false, "do not use sparse index")
	sparseCheckout_initCmd.Flags().Bool("sparse-index", false, "use sparse index")
	sparseCheckoutCmd.AddCommand(sparseCheckout_initCmd)
}

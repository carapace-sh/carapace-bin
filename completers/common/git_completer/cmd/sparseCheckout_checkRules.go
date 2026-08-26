package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_checkRulesCmd = &cobra.Command{
	Use:   "check-rules",
	Short: "Check whether sparsity rules match one or more paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_checkRulesCmd).Standalone()

	sparseCheckout_checkRulesCmd.Flags().Bool("no-cone", false, "use non-cone mode")
	sparseCheckout_checkRulesCmd.Flags().BoolP("null", "z", false, "terminate input and output paths with NUL byte")
	sparseCheckoutCmd.AddCommand(sparseCheckout_checkRulesCmd)
}

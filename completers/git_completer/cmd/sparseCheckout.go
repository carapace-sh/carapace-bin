package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckoutCmd = &cobra.Command{
	Use:     "sparse-checkout",
	Short:   "Initialize and modify the sparse-checkout",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(sparseCheckoutCmd).Standalone()

	rootCmd.AddCommand(sparseCheckoutCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove files outside the sparse-checkout definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_cleanCmd).Standalone()

	sparseCheckout_cleanCmd.Flags().Bool("dry-run", false, "list directories that would be removed without deleting them")
	sparseCheckout_cleanCmd.Flags().BoolP("force", "f", false, "allow deletion of files outside the sparse-checkout definition")
	sparseCheckout_cleanCmd.Flags().Bool("verbose", false, "list every file within the directories considered for removal")
	sparseCheckoutCmd.AddCommand(sparseCheckout_cleanCmd)
}

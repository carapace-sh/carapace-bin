package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCmd = &cobra.Command{
	Use:   "sparse",
	Short: "Manage which paths from the working-copy commit are present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCmd).Standalone()

	sparseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(sparseCmd)
}

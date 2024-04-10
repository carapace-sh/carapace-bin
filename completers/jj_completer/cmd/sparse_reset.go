package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparse_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the patterns to include all files in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_resetCmd).Standalone()

	sparse_resetCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	sparseCmd.AddCommand(sparse_resetCmd)
}

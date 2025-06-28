package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparse_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Start an editor to update the patterns that are present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_editCmd).Standalone()

	sparse_editCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	sparseCmd.AddCommand(sparse_editCmd)
}

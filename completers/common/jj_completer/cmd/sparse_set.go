package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparse_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update the patterns that are present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_setCmd).Standalone()

	sparse_setCmd.Flags().StringSlice("add", nil, "Patterns to add to the working copy")
	sparse_setCmd.Flags().Bool("clear", false, "Include no files in the working copy (combine with --add)")
	sparse_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	sparse_setCmd.Flags().StringSlice("remove", nil, "Patterns to remove from the working copy")
	sparseCmd.AddCommand(sparse_setCmd)
}

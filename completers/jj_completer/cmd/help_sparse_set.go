package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_sparse_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update the patterns that are present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sparse_setCmd).Standalone()

	help_sparseCmd.AddCommand(help_sparse_setCmd)
}

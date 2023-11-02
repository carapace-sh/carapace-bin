package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_sparse_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the patterns that are currently present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sparse_listCmd).Standalone()

	help_sparseCmd.AddCommand(help_sparse_listCmd)
}

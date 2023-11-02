package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparse_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the patterns that are currently present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_help_listCmd).Standalone()

	sparse_helpCmd.AddCommand(sparse_help_listCmd)
}

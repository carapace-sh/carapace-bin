package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparse_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update the patterns that are present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_help_setCmd).Standalone()

	sparse_helpCmd.AddCommand(sparse_help_setCmd)
}

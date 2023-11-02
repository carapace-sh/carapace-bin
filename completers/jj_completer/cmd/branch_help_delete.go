package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing branch and propagate the deletion to remotes on the next push",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_deleteCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_deleteCmd)
}

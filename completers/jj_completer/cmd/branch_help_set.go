package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a given branch to point to a certain commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_setCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_setCmd)
}

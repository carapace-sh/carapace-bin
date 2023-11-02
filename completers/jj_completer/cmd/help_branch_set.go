package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branch_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a given branch to point to a certain commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_setCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_setCmd)
}

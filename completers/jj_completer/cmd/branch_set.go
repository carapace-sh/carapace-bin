package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Update a given branch to point to a certain commit",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_setCmd).Standalone()

	branch_setCmd.Flags().BoolP("allow-backwards", "B", false, "Allow moving the branch backwards or sideways")
	branch_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_setCmd.Flags().StringP("revision", "r", "", "The branch's target revision")
	branchCmd.AddCommand(branch_setCmd)
}

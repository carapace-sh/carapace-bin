package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_deleteCmd = &cobra.Command{
	Use:     "delete [OPTIONS] [NAMES]...",
	Short:   "Delete an existing branch and propagate the deletion to remotes on the next push",
	Aliases: []string{"d"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_deleteCmd).Standalone()

	branch_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_deleteCmd)

	carapace.Gen(branch_deleteCmd).PositionalAnyCompletion(
		jj.ActionLocalBranches().FilterArgs(),
	)
}

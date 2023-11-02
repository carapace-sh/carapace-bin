package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_forgetCmd = &cobra.Command{
	Use:     "forget [OPTIONS] [NAMES]...",
	Short:   "Forget everything about a branch, including its local and remote targets",
	Aliases: []string{"f"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_forgetCmd).Standalone()

	branch_forgetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_forgetCmd)

	carapace.Gen(branch_forgetCmd).PositionalAnyCompletion(
		jj.ActionLocalBranches().FilterArgs(),
	)
}

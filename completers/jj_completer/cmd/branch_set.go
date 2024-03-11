package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_setCmd = &cobra.Command{
	Use:     "set [OPTIONS] <NAMES>...",
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

	carapace.Gen(branch_setCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(branch_setCmd).PositionalAnyCompletion(
		jj.ActionLocalBranches().FilterArgs(),
	)
}

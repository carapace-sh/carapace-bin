package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_moveCmd = &cobra.Command{
	Use:     "move [OPTIONS] <--from <REVISIONS>|NAME>",
	Short:   "Move existing branches to target revision",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_moveCmd).Standalone()

	branch_moveCmd.Flags().BoolP("allow-backwards", "B", false, "Allow moving the branch backwards or sideways")
	branch_moveCmd.Flags().String("from", "@", "Move part of this change into the destination")
	branch_moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_moveCmd.Flags().String("to", "@", "Move part of the source into this change")
	branchCmd.AddCommand(branch_moveCmd)

	carapace.Gen(branch_moveCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(branch_moveCmd).PositionalAnyCompletion(
		jj.ActionLocalBranches().FilterArgs(),
	)
}

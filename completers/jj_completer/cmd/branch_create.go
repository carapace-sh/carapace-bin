package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_createCmd = &cobra.Command{
	Use:     "create [OPTIONS] <NAMES>...",
	Short:   "Create a new branch",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_createCmd).Standalone()

	branch_createCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_createCmd.Flags().StringP("revision", "r", "", "The branch's target revision")
	branchCmd.AddCommand(branch_createCmd)

	carapace.Gen(branch_createCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevsOption{}.Default()),
	})

	carapace.Gen(branch_createCmd).PositionalAnyCompletion(
		jj.ActionLocalBranches().FilterArgs(),
	)
}

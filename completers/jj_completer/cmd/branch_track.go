package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var branch_trackCmd = &cobra.Command{
	Use:   "track [OPTIONS] <NAMES>...",
	Short: "Start tracking given remote branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_trackCmd).Standalone()

	branch_trackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_trackCmd)

	carapace.Gen(branch_trackCmd).PositionalAnyCompletion(
		jj.ActionRemoteBranches("").FilterArgs(),
	)
}

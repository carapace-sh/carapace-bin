package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var step_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Fast-forward target to current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_pushCmd).Standalone()

	step_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_pushCmd)

	carapace.Gen(step_pushCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}

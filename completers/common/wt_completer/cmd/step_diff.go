package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var step_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show all changes since branching",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_diffCmd).Standalone()

	step_diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_diffCmd)

	carapace.Gen(step_diffCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(step_diffCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := step_diffCmd.Flags().Args()
			if len(args) == 0 {
				return carapace.ActionValues()
			}
			return bridge.ActionCarapaceBin("git", "diff", args[0])
		}),
	)
}

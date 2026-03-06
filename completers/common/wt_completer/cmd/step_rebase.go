package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var step_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Rebase onto target",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_rebaseCmd).Standalone()

	step_rebaseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_rebaseCmd)

	carapace.Gen(step_rebaseCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}

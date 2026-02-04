package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a commit to a different location in the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	moveCmd.Flags().BoolP("after", "a", false, "Move the commit after (above) the target instead of before (below)")
	moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(moveCmd)

	carapace.Gen(moveCmd).FlagCompletion(carapace.ActionMap{
		"after": but.ActionTargets(),
	})

	carapace.Gen(moveCmd).PositionalCompletion(
		but.ActionCommits(),
		but.ActionTargets(),
	)
}

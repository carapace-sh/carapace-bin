package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a commit or branch to a different location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	moveCmd.Flags().BoolP("after", "a", false, "Move the commit after (above) the target instead of before (below). Only valid for commit-to-commit moves")
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

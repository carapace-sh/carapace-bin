package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var commit_emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Insert a blank commit before or after the specified commit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commit_emptyCmd).Standalone()

	commit_emptyCmd.Flags().String("after", "", "Insert the blank commit after this commit or branch")
	commit_emptyCmd.Flags().String("before", "", "Insert the blank commit before this commit or branch")
	commit_emptyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.AddCommand(commit_emptyCmd)

	carapace.Gen(commit_emptyCmd).FlagCompletion(carapace.ActionMap{
		"after":  but.ActionTargets(),
		"before": but.ActionTargets(),
	})

	carapace.Gen(commit_emptyCmd).PositionalCompletion(
		but.ActionTargets(),
	)
}

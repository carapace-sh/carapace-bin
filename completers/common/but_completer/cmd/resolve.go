package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "Resolve conflicts in a commit",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().Bool("ai", false, "Resolve the conflicts with the configured AI model and apply the result")
	resolveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(resolveCmd)

	carapace.Gen(resolveCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	)
}

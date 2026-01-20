package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var uncommitCmd = &cobra.Command{
	Use:   "uncommit",
	Short: "Uncommit changes from a commit or file-in-commit to the unstaged area",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uncommitCmd).Standalone()

	uncommitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(uncommitCmd)

	carapace.Gen(uncommitCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	)
}

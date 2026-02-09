package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows detailed information about a commit or branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	showCmd.Flags().BoolP("verbose", "v", false, "Show full commit messages and files changed for each commit")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionTargets(),
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Commits: true}),
		).ToA(),
	)
}

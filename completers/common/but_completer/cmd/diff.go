package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Displays the diff of changes in the repo.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffCmd.Flags().Bool("no-tui", false, "Disable the interactive TUI diff viewer (overrides but.ui.tui config)")
	diffCmd.Flags().Bool("tui", false, "Open an interactive TUI diff viewer")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionChanges(git.ChangeOpts{}.Default()),
			but.ActionTargets(),
			but.ActionCliIds(but.CliIdsOpts{}.Default()),
		).ToA(),
	)
}

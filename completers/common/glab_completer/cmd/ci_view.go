package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_viewCmd = &cobra.Command{
	Use:   "view [branch/tag]",
	Short: "View, run, trace, log, and cancel CI/CD job's current pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_viewCmd).Standalone()

	ci_viewCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch or tag. Defaults to the current branch.")
	ci_viewCmd.Flags().BoolP("web", "w", false, "Open pipeline in a browser. Uses default browser, or browser specified in BROWSER variable.")
	ciCmd.AddCommand(ci_viewCmd)

	carapace.Gen(ci_viewCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_viewCmd),
	})

	carapace.Gen(ci_viewCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionBranches(ci_viewCmd),
			action.ActionReleases(ci_viewCmd),
		).ToA(),
	)
}

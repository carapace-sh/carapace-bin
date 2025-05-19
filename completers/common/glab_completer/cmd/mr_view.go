package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_viewCmd = &cobra.Command{
	Use:     "view {<id> | <branch>}",
	Short:   "Display the title, body, and other information about a merge request.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_viewCmd).Standalone()

	mr_viewCmd.Flags().BoolP("comments", "c", false, "Show merge request comments and activities.")
	mr_viewCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	mr_viewCmd.Flags().StringP("page", "p", "", "Page number.")
	mr_viewCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	mr_viewCmd.Flags().BoolP("system-logs", "s", false, "Show system activities and logs.")
	mr_viewCmd.Flags().BoolP("web", "w", false, "Open merge request in a browser. Uses default browser or browser specified in BROWSER variable.")
	mrCmd.AddCommand(mr_viewCmd)

	carapace.Gen(mr_viewCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})

	carapace.Gen(mr_viewCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_viewCmd, "opened"),
	)
}

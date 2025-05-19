package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_viewCmd = &cobra.Command{
	Use:     "view <id>",
	Short:   "Display the title, body, and other information about an issue.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_viewCmd).Standalone()

	issue_viewCmd.Flags().BoolP("comments", "c", false, "Show issue comments and activities.")
	issue_viewCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	issue_viewCmd.Flags().StringP("page", "p", "", "Page number.")
	issue_viewCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	issue_viewCmd.Flags().BoolP("system-logs", "s", false, "Show system activities and logs.")
	issue_viewCmd.Flags().BoolP("web", "w", false, "Open issue in a browser. Uses the default browser, or the browser specified in the $BROWSER variable.")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(issue_closeCmd, "opened"),
	)
}

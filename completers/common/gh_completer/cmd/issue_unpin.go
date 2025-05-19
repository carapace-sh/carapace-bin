package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_unpinCmd = &cobra.Command{
	Use:     "unpin {<number> | <url>}",
	Short:   "Unpin a issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_unpinCmd).Standalone()

	issueCmd.AddCommand(issue_unpinCmd)

	carapace.Gen(issue_unpinCmd).PositionalCompletion(
		action.ActionPinnedIssues(issue_unpinCmd),
	)
}

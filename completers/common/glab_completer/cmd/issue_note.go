package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_noteCmd = &cobra.Command{
	Use:     "note <issue-id>",
	Short:   "Comment on an issue in GitLab.",
	Aliases: []string{"comment"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_noteCmd).Standalone()

	issue_noteCmd.Flags().StringP("message", "m", "", "Message text.")
	issueCmd.AddCommand(issue_noteCmd)

	carapace.Gen(issue_noteCmd).PositionalCompletion(
		action.ActionIssues(issue_noteCmd, "opened"),
	)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_artifact_deleteCmd = &cobra.Command{
	Use:   "delete {<issue> <artifact-number>... | <artifact-url>...}",
	Short: "Delete issue artifacts (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifact_deleteCmd).Standalone()

	issue_artifact_deleteCmd.Flags().Bool("yes", false, "Confirm deletion without prompting")
	issue_artifactCmd.AddCommand(issue_artifact_deleteCmd)

	carapace.Gen(issue_artifact_deleteCmd).PositionalCompletion(
		action.ActionIssues(issue_artifact_deleteCmd, action.IssueOpts{Open: true, Closed: true}),
	)

	carapace.Gen(issue_artifact_deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionIssueArtifacts(issue_artifact_deleteCmd, c.Args[0])
		}),
	)
}

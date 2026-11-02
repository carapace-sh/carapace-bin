package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_artifact_editCmd = &cobra.Command{
	Use:   "edit {<issue> <artifact-number> | <artifact-url>}",
	Short: "Edit an issue artifact (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifact_editCmd).Standalone()

	issue_artifact_editCmd.Flags().StringSlice("attach", nil, "Attach an image or video `file`, in '<file>#<image alt text>' format")
	issue_artifact_editCmd.Flags().StringP("body", "b", "", "New content; for a link, a URL")
	issue_artifact_editCmd.Flags().StringP("body-file", "F", "", "Read new content from `file` (use \"-\" to read from standard input)")
	issue_artifact_editCmd.Flags().String("name", "", "New name")
	issue_artifact_editCmd.Flags().String("restore-version", "", "Restore the version with this `number` from the edit history")
	issue_artifactCmd.AddCommand(issue_artifact_editCmd)

	carapace.Gen(issue_artifact_editCmd).FlagCompletion(carapace.ActionMap{
		"attach":    carapace.ActionFiles(),
		"body":      action.ActionBody(issue_artifact_editCmd),
		"body-file": carapace.ActionFiles(),
		"restore-version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return action.ActionIssueArtifactVersions(issue_artifact_editCmd, c.Args[0], c.Args[1])
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(issue_artifact_editCmd).PositionalCompletion(
		action.ActionIssues(issue_artifact_editCmd, action.IssueOpts{Open: true, Closed: true}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionIssueArtifacts(issue_artifact_editCmd, c.Args[0])
		}),
	)
}

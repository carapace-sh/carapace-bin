package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var issue_artifact_viewCmd = &cobra.Command{
	Use:   "view {<issue> <artifact-number> | <artifact-url>}",
	Short: "View an issue artifact (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifact_viewCmd).Standalone()

	issue_artifact_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_artifact_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	issue_artifact_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	issue_artifact_viewCmd.Flags().String("version", "", "Show the version with this `number` from the edit history")
	issue_artifact_viewCmd.Flags().BoolP("web", "w", false, "Open a document's page on GitHub, or a link's URL")
	issue_artifactCmd.AddCommand(issue_artifact_viewCmd)

	carapace.Gen(issue_artifact_viewCmd).FlagCompletion(carapace.ActionMap{
		"jq":   jq.ActionFilters(),
		"json": carapace.ActionValues("body", "bodyHtml", "createdAt", "creator", "description", "id", "name", "number", "type", "updatedAt", "updatedByActor", "url", "versions").UniqueList(","),
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return action.ActionIssueArtifactVersions(issue_artifact_viewCmd, c.Args[0], c.Args[1])
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(issue_artifact_viewCmd).PositionalCompletion(
		action.ActionIssues(issue_artifact_viewCmd, action.IssueOpts{Open: true, Closed: true}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionIssueArtifacts(issue_artifact_viewCmd, c.Args[0])
		}),
	)
}

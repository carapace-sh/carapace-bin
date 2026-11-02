package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var issue_artifact_listCmd = &cobra.Command{
	Use:     "list {<issue-number> | <issue-url>}",
	Short:   "List artifacts on an issue (preview)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifact_listCmd).Standalone()

	issue_artifact_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_artifact_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	issue_artifact_listCmd.Flags().StringP("limit", "L", "", "Maximum number of artifacts to fetch")
	issue_artifact_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	issue_artifact_listCmd.Flags().String("type", "", "Filter by type: {generic|plan|link|embedded_html}")
	issue_artifactCmd.AddCommand(issue_artifact_listCmd)

	carapace.Gen(issue_artifact_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":   jq.ActionFilters(),
		"json": carapace.ActionValues("body", "bodyHtml", "createdAt", "creator", "description", "id", "name", "number", "type", "updatedAt", "updatedByActor", "url").UniqueList(","),
		"type": carapace.ActionValues("embedded_html", "generic", "link", "plan"),
	})

	carapace.Gen(issue_artifact_listCmd).PositionalCompletion(
		action.ActionIssues(issue_artifact_listCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}

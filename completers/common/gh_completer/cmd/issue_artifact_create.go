package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_artifact_createCmd = &cobra.Command{
	Use:     "create {<issue-number> | <issue-url>} [<file>[#<name>] | <url>]...",
	Short:   "Create issue artifacts (preview)",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifact_createCmd).Standalone()

	issue_artifact_createCmd.Flags().StringSlice("attach", nil, "Attach an image or video `file`, in '<file>#<image alt text>' format")
	issue_artifact_createCmd.Flags().StringP("body", "b", "", "Content for a single document")
	issue_artifact_createCmd.Flags().StringP("body-file", "F", "", "Read content for a single artifact from `file` (use \"-\" to read from standard input)")
	issue_artifact_createCmd.Flags().String("name", "", "Name for a single artifact")
	issue_artifact_createCmd.Flags().String("type", "", "Type for documents: {generic|plan|embedded_html}")
	issue_artifactCmd.AddCommand(issue_artifact_createCmd)

	carapace.Gen(issue_artifact_createCmd).FlagCompletion(carapace.ActionMap{
		"attach":    carapace.ActionFiles(),
		"body":      action.ActionBody(issue_artifact_createCmd),
		"body-file": carapace.ActionFiles(),
		"type":      carapace.ActionValues("embedded_html", "generic", "plan"),
	})

	carapace.Gen(issue_artifact_createCmd).PositionalCompletion(
		action.ActionIssues(issue_artifact_createCmd, action.IssueOpts{Open: true, Closed: true}),
	)

	carapace.Gen(issue_artifact_createCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}

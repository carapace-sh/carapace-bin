package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show status of relevant issues",
	GroupID: "General commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_statusCmd).Standalone()

	issue_statusCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_statusCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	issue_statusCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	issueCmd.AddCommand(issue_statusCmd)

	carapace.Gen(issue_statusCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionIssueFields().UniqueList(","),
	})
}

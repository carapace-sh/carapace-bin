package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show status of relevant pull requests",
	GroupID: "General commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_statusCmd).Standalone()

	pr_statusCmd.Flags().BoolP("conflict-status", "c", false, "Display the merge conflict status of each pull request")
	pr_statusCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_statusCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	pr_statusCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	prCmd.AddCommand(pr_statusCmd)

	carapace.Gen(pr_statusCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionPullRequestFields().UniqueList(","),
	})
}

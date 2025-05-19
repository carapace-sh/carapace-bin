package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_boardCmd = &cobra.Command{
	Use:   "board [command] [flags]",
	Short: "Work with GitLab Issue Boards in the given project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_boardCmd).Standalone()

	issue_boardCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the OWNER/REPO format or the project ID. Supports group namespaces")
	issueCmd.AddCommand(issue_boardCmd)

	carapace.Gen(issue_boardCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(issue_boardCmd),
	})
}

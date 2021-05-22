package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workflowCmd = &cobra.Command{
	Use:   "workflow <command>",
	Short: "View details about GitHub Actions workflows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	workflowCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(workflowCmd)

	carapace.Gen(workflowCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionOwnerRepositories(workflowCmd),
	})
}

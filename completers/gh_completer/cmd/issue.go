package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue <command>",
	Short: "Manage issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issueCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionOwnerRepositories(issueCmd),
	})
}

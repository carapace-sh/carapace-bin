package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:     "issue <command>",
	Short:   "Manage issues",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issueCmd).Standalone()
	issueCmd.AddGroup(
		&cobra.Group{ID: "general", Title: "General commands"},
		&cobra.Group{ID: "targeted", Title: "Targeted commands"},
	)

	issueCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepoOverride(issueCmd),
	})
}
